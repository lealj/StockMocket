package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/gorilla/securecookie"
	"gorm.io/gorm"
)

var (
	jwtSecretKey = []byte("thekeythatwillbeused")
	securityKey  = securecookie.New([]byte("authorization key"), nil)
)

type Credentials struct {
	gorm.Model

	Username string  `json:"username"`
	Password string  `json:"password"`
	Funds    float32 `json:"funds"`
}

func login(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	// userCredentials has two field user and pass
	var userCredentials = &Credentials{}
	err := json.NewDecoder(router.Body).Decode(&userCredentials)
	// This is a precaution so body is correct
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	existingCredentials := &Credentials{}
	// Getting users with same username that are already in the database
	DB.Table("credentials").Select("username", "password").Where("username = ?", userCredentials.Username).Scan(&existingCredentials)

	if err != nil {
		writer.WriteHeader(http.StatusBadGateway)
		return
	}
	if err != nil {
		// if no user with the same username, send back code 401
		if err == sql.ErrNoRows {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}
		// if error is one of any other type, send 500 status
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	// if the passwords don't match, we send back another 401
	if existingCredentials.Password != userCredentials.Password {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	// When we get here, that means that the passwords matched

	// JWT Token is created if user has an account.
	token, err := generateToken(userCredentials.Username)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(writer, &http.Cookie{
		Name:     "loggedIn",
		Value:    token,
		HttpOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(24 * time.Hour),
	})

	//log.Printf(token)
	writer.WriteHeader(http.StatusOK)
	//log.Printf("The passwords matched, the status code should be 200\n")

}

func signup(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var newCredentials Credentials
	//default variable value
	newCredentials.Funds = 1000
	//handles data received from request (json data)
	err := json.NewDecoder(router.Body).Decode(&newCredentials)

	// This is a precaution so body is correct
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	existingCredentials := &Credentials{}
	// Getting users with same username that are already in the database
	DB.Table("credentials").Select("username", "password").Where("username = ?", newCredentials.Username).Scan(&existingCredentials)
	//check if new signup username already exists
	if existingCredentials.Username == newCredentials.Username || existingCredentials.Username != "" {
		log.Println("Username already exists")
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	//create new user as username doesn't exist already

	DB.Create(&newCredentials)

	err = json.NewEncoder(writer).Encode(newCredentials)
	if err != nil {
		return
	}
	writer.WriteHeader(http.StatusOK)
	log.Printf("Successfully saved username and password")
}

func deleteCredentials(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	// Get the username from the body to delete and checks to see body is correct
	var credentialsToDelete Credentials
	err := json.NewDecoder(router.Body).Decode(&credentialsToDelete)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	// Now we must check if the username is in the database
	var existingCredentials Credentials
	if err := DB.Table("credentials").Select("username", "password").Where("username = ?",
		credentialsToDelete.Username).Scan(&existingCredentials); err != nil {
	}

	// If the username exists, we delete it and send a 200 code to suggest we deleted successfully. Else we return 409
	if existingCredentials.Username == credentialsToDelete.Username || existingCredentials.Username != "" {
		log.Println("Username to delete is in database")

		if err := DB.Where("username = ?", credentialsToDelete.Username).Unscoped().Delete(&existingCredentials); err.Error != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		log.Println("Deleted in DB")
		writer.WriteHeader(http.StatusOK)
		return

	} else {
		writer.WriteHeader(http.StatusConflict)
	}

}

// Logout might have to get changed to be path protected. There is no reason for someone to attempt to log out if they are not
// even logged in the first place. This requires changes to the main.go file to make this go through the JWTHandler.
// Should be discussed first.
func logout(w http.ResponseWriter, r *http.Request) {
	// deletes cookies from local browser storage
	http.SetCookie(w, &http.Cookie{
		Name:     "loggedIn",
		Value:    "",
		HttpOnly: true,
		Secure:   true,
		Expires:  time.Unix(0, 0),
	})

	w.WriteHeader(http.StatusOK)
}

// GetUserFunds is protected and user must be logged in to get the user funds
func GetUserFunds(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	// instead of having to send the username each time, as long as the user is logged in, the HTTPCookie can gather
	// the username from the claims. All you must do is call it from the front end with an empty json body
	// It should work as of right now with no changed to other sections but more code can be done like this now.
	cookie, err := router.Cookie("loggedIn")
	if err != nil {
		if err == http.ErrNoCookie {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	tokenStr := cookie.Value

	claims, err := extractClaims(tokenStr)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	// should jsut contain username
	var creds Credentials

	if err != nil {
		fmt.Printf("Error decoding: %v", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	// Now we must check if the username is in the database
	if err := DB.Where("username = ?", claims.Username).First(&creds).Error; err != nil {
		fmt.Printf("Error finding username\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("Funds Test %v", creds.Funds)

	err = json.NewEncoder(writer).Encode(&creds.Funds)
	if err != nil {
		return
	}
}
