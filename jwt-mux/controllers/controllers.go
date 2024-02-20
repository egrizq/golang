package controllers

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
	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.Users

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error parsing body: %s", err)
		log.Printf("error parsing body: %s", err)
		return
	}
	defer r.Body.Close()

	// todo hash password
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashPassword)

	query := "INSERT INTO mydata(id, fullname, email, password) VALUES ($1, $2, $3, $4)"
	_, err = database.DB.Exec(query, user.Id, user.Fullname, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error inserting data: %s", err)
		log.Printf("Error inserting data: %s", err)
		return
	}

	fmt.Fprintf(w, "fullname: %s & email: %s", user.Fullname, user.Email)
	log.Println("Register success!")
}

func Login(w http.ResponseWriter, r *http.Request) {
	var userInput models.UserInput

	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error parsing body: %s", err)
		log.Printf("error parsing body: %s", err)
		return
	}
	defer r.Body.Close()
	log.Println(userInput.Email, userInput.Password)

	query := "SELECT password FROM mydata WHERE email = $1"
	rows, err := database.DB.Query(query, userInput.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "No data in database: %s", err)
		log.Printf("No data in database: %s", err)
		return
	}
	defer rows.Close()

	var password string
	if rows.Next() {
		if err := rows.Scan(&password); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "cannot get password: %s", err)
			return
		}
	} else {
		// No rows returned
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "No data found in database")
		log.Println("No data found in database")
		return
	}
	log.Println("password:", password)

	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(userInput.Password)); err != nil {
		w.WriteHeader(http.StatusBadGateway)
		fmt.Printf("password is incorrect: %s", err)
		return
	}

	// todo jwt
	expTime := time.Now().Add(time.Second * 60)
	claims := &config.JWTClaim{
		Email: userInput.Email,
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
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error in token: %s", err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})

	fmt.Fprintf(w, "Login success!")
	log.Println("Login success!")
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})

	fmt.Fprintf(w, "Logout success!")
	log.Println("Logout success!")
}

func Alldata(w http.ResponseWriter, r *http.Request) {
	query := "SELECT fullname, email  FROM mydata"
	rows, err := database.DB.Query(query)
	if err != nil {
		log.Printf("error query: %s", err.Error())
		fmt.Fprintf(w, "error query: %s", err.Error())
		return
	}

	var allUser []models.StoreUser

	for rows.Next() {
		var dataUser models.StoreUser

		if err := rows.Scan(&dataUser.Fullname, &dataUser.Email); err != nil {
			log.Printf("error scan data: %s", err.Error())
			fmt.Fprintf(w, "error scan data: %s", err.Error())
			return
		}

		allUser = append(allUser, dataUser)
	}

	log.Println(allUser)
	jsonData, err := json.Marshal(allUser)
	if err != nil {
		log.Printf("error turn data to json: %s", err.Error())
		fmt.Fprintf(w, "error turn data to json: %s", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonData); err != nil {
		log.Printf("error app: %s", err.Error())
		fmt.Fprintf(w, "error app: %s", err.Error())
		return
	}
}
