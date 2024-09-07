package handlers

import (
    "database/sql"
    "encoding/json"
    "net/http"
    "strconv"
	"product-service/models"

    _ "github.com/lib/pq"
    "github.com/gorilla/mux"
)

var db *sql.DB

func init() {
    var err error
    connStr := "user=postgres password=mysecretpassword dbname=food_delivery sslmode=disable"
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        panic(err)
    }
}

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("SELECT id, name, description, price, image_url FROM products")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var products []models.Products
    for rows.Next() {
        var product models.Products
        if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.ImageUrl); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        products = append(products, product)
    }

    json.NewEncoder(w).Encode(products)
}

func GetProductByIDHandler(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    var product models.Products
    err := db.QueryRow("SELECT id, name, description, price, image_url FROM products WHERE id=$1", id).Scan(
        &product.ID, &product.Name, &product.Description, &product.Price, &product.ImageUrl)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(product)
}
