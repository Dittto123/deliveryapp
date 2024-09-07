package main

import (
	"log"
	"net/http"
	"order-service/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/orders", handlers.CreateOrderHandler).Methods("POST")
	r.HandleFunc("/orders/{id}", handlers.GetOrderByIdHandler).Methods("GET")
	r.HandleFunc("/orders/{id}/status", handlers.UpdateOrderHandler).Methods("PATCH")
	
	
	log.Fatal(http.ListenAndServe(":8002", r))
}