package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

type MyClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`

	jwt.StandardClaims
}

// JWTPathProtection This handler can be applied to anything that requires the user to log in. To view an example of how to do this. View
// main.go and the comment starting in A1
func JWTPathProtection(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("loggedIn")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
			if jwt.GetSigningMethod("HS256") != token.Method {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Method)
			}
			return jwtSecretKey, nil
		})

		if err != nil || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusOK)

		// Call the next handler after token is verified
		next.ServeHTTP(w, r)
	})
}

func getClaimsHandler(w http.ResponseWriter, r *http.Request) {
	// This is a handler to get claims. Claims can be obtained within any request that is asked from frontend as long as with
	// credentials is true in the request. Example of this in claimData in frontend login service
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
	claims, err := extractClaims(tokenStr)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// We set the status code as good and then return all the valid information.
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(claims)
	if err != nil {
		return
	}
}

func generateToken(username string) (string, error) {
	// Generates a token that stores encoded information using HS256 bit encryption. This generates a token that is used
	// with HTTPOnly cookies. This means that the cookie is transferred safely and is more protected from XSS attacks.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyClaims{
		Username: username,
		Role:     "admin",
		StandardClaims: jwt.StandardClaims{
			Issuer:    "StockMocket",
			Subject:   "TheCODE",
			Audience:  username,
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
			Id:        "33",
		},
	})

	return token.SignedString(jwtSecretKey)
}

func extractClaims(tokenString string) (*MyClaims, error) {
	// This works in conjunction with getClaimsHandler. This will return claims as well as verify the token. If the user is
	// not logged in, this will verify that.
	// This can be called to extract claims from anywhere in the DB.
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*MyClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
