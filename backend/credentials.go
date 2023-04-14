package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/securecookie"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

var (
	jwtSecretKey = []byte("thekeythatwillbeused")
	securityKey  = securecookie.New([]byte("authorization key"), nil)
)

type Credentials struct {
	gorm.Model

	Username string  `json:"username"`
	Password string  `json:"password"`
	Funds    float64 `json:"funds"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
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

	log.Printf(token)
	writer.WriteHeader(http.StatusOK)
	log.Printf("The passwords matched, the status code should be 200\n")

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
	log.Printf(newCredentials.Username)

	DB.Create(&newCredentials)
	json.NewEncoder(writer).Encode(newCredentials)

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

func verifyToken(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("loggedIn")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := cookie.Value
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	log.Printf(tokenStr)

	json.NewEncoder(w).Encode(claims)
}

func ParseToken(tokenStr string) (*jwt.StandardClaims, error) {
	claims := &jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}

func generateToken(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.StandardClaims{
		Id:        username,
		ExpiresAt: expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecretKey)
}

func authenticateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		claims, err := ParseToken(cookie.Value)
		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		// Token is valid, pass it to the next middleware or handler
		ctx := context.WithValue(r.Context(), "username", claims.Id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func logout(w http.ResponseWriter, r *http.Request) {
	// delete JWT cookie

	http.SetCookie(w, &http.Cookie{
		Name:     "loggedIn",
		Value:    "",
		HttpOnly: true,
		Secure:   true,
		Expires:  time.Unix(0, 0),
	})

	w.WriteHeader(http.StatusOK)
}
