package main

import (
	"fmt"
	"log"
	"net/http"
	"product-service/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/products", handlers.GetProductsHandler).Methods("GET")
	r.HandleFunc("/products/{id}", handlers.GetProductByIDHandler).Methods("GET")
	
	fmt.Printf("hello")
	log.Fatal(http.ListenAndServe(":8001", r))
}