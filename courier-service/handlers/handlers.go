package handlers

import (
    "database/sql"
    "encoding/json"
    "net/http"
    "strconv"
    "github.com/gorilla/mux"
    _ "github.com/lib/pq"
	"courier-service/models"
)

var db *sql.DB

func init() {
    var err error
    connStr := "user=postgres password=diyor938 dbname=food_delivery sslmode=disable"
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        panic(err)
    }
}

func GetAssignedOrdersHandler(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    courierID, _ := strconv.Atoi(params["id"])

    rows, err := db.Query("SELECT id, product_id, quantity, status FROM orders WHERE courier_id=$1", courierID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var orders []models.Order
    for rows.Next() {
        var order models.Order
        if err := rows.Scan(&order.ID, &order.ProductID, &order.Quantity, &order.Status); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        orders = append(orders, order)
    }

    json.NewEncoder(w).Encode(orders)
}

func UpdateOrderStatusHandler(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    orderID, _ := strconv.Atoi(params["order_id"])

    var statusUpdate struct {
        Status string `json:"status"`
    }
    json.NewDecoder(r.Body).Decode(&statusUpdate)

    _, err := db.Exec("UPDATE orders SET status=$1 WHERE id=$2", statusUpdate.Status, orderID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
