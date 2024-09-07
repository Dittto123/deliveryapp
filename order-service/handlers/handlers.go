package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"order-service/models"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB

func Init() {
	var err error
	conStr := "user=postgres password=diyor938 dbname=postgres sslmode=disable"
	db, err = sql.Open("postgres", conStr)
	if err != nil {
		panic(err)
	}
}

func CreateOrderHandler(res http.ResponseWriter, req *http.Request) {
	var order models.Order
    json.NewDecoder(req.Body).Decode(&order)

    query := "INSERT INTO orders (user_id, product_id, quantity, status) VALUES ($1, $2, $3, 'pending') RETURNING id"
    err := db.QueryRow(query, order.UserID, order.ProductID, order.Quantity).Scan(&order.ID)
    if err != nil {
        http.Error(res, err.Error(), http.StatusInternalServerError)
        return
    }

	log.Fatal(order.ProductID)
    json.NewEncoder(res).Encode(order)
}

func GetOrderByIdHandler(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
    id, _ := strconv.Atoi(params["id"])

    var order models.Order
    err := db.QueryRow("SELECT id, user_id, product_id, quantity, status FROM orders WHERE id=$1", id).Scan(
        &order.ID, &order.UserID, &order.ProductID, &order.Quantity, &order.Status)
    if err != nil {
        http.Error(res, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(res).Encode(order)
}

func UpdateOrderHandler(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
    id, _ := strconv.Atoi(params["id"])

    var statusUpdate struct {
        Status string `json:"status"`
    }
    json.NewDecoder(req.Body).Decode(&statusUpdate)

    _, err := db.Exec("UPDATE orders SET status=$1 WHERE id=$2", statusUpdate.Status, id)
    if err != nil {
        http.Error(res, err.Error(), http.StatusInternalServerError)
        return
    }

    res.WriteHeader(http.StatusNoContent)
}