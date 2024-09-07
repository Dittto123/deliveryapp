package main

import (
	"courier-service/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/couriers/{id}/orders", handlers.GetAssignedOrdersHandler).Methods("GET")
	r.HandleFunc("couriers/{id}/orders/{order_id}/status", handlers.UpdateOrderStatusHandler).Methods("PATCH")

	fmt.Println("server started")
	log.Fatal(http.ListenAndServe(":8003", r))
}