package handlers

import (
	"auth-service/models"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var dataBS *sql.DB

func InitDB() {
	var err error
	conStr := "user=postgres password=diyor938 dbname=postgres sslmode=disable"
	dataBS, err = sql.Open("postgres", conStr)
	if err != nil {
		panic(err)
	} 
}

func SignUpHandler(res http.ResponseWriter, req *http.Request) {
	var user models.User
	json.NewDecoder(req.Body).Decode(&user)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	_, err := dataBS.Exec("INSERT INTO users (email, password, role) VALUES($1, $2, $3) RETURNING ID", user.Email, user.Password, user.Role)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Sign up attempt: Email: %s, Password: %s", user.Email, user.Password)

	res.WriteHeader(http.StatusCreated)
}

func LoginHandler(res http.ResponseWriter, req *http.Request) {
	var acc models.Accounts
	json.NewDecoder(req.Body).Decode(&acc)

	token, err := generateGWT(acc.Email)
	if err != nil {
		http.Error(res, "Error generating token", http.StatusInternalServerError)
		return
	}

	log.Printf("Email: %s, Password: %s", acc.Email, acc.Password)

	json.NewEncoder(res).Encode(map[string]string{"token": token})
}

var jwtKey = "secret key"

func generateGWT(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"Exp":   time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtKey))
}
