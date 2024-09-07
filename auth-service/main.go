package main

import (
	"log"
	"net/http"
	"auth-service/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/auth/login", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/auth/signUp", handlers.SignUpHandler).Methods("POST")

	log.Println("server started")
	log.Fatal(http.ListenAndServe(":8000", r))
}
