package helper

import (
	"encoding/json"
	"fmt"
	"jwt-mux/config"
	"jwt-mux/database"
	"jwt-mux/models"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func RegisterData(w http.ResponseWriter, query string, id int, fullname string, email string, password string) {
	_, err := database.DB.Exec(query, id, fullname, email, password)
	if err != nil {
		fmt.Fprintf(w, "Error inserting data: %s", err)
		return
	}

	log.Println("Register success!")
}

func DataToJson(w http.ResponseWriter, allUser []models.StoreUser) {
	jsonData, err := json.Marshal(allUser)
	if err != nil {
		fmt.Fprintf(w, "error turn data to json: %s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonData); err != nil {
		fmt.Fprintf(w, "error set to JSON: %s", err)
		return
	}
}

func JwtGenerate(w http.ResponseWriter, email string) {

	// add expired time
	expTime := time.Now().Add(time.Second * 60)

	claims := &config.JWTClaim{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "golang-otp",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	// declare tokens
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// signed token
	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		fmt.Fprintf(w, "Error in token: %s", err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})
}
