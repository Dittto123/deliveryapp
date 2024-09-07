package main

import (
	"fmt"
	"log"
	"net/http"
	"notification-service/handlers"

	"github.com/gorilla/mux"
)

func main () {
	r := mux.NewRouter()
	r.HandleFunc("/notifications/send", handlers.SendNotification).Methods("POST")

	fmt.Println("server-started")
	log.Fatal(http.ListenAndServe(":8005", r))
}