package main

import (
	"encoding/json"
	"net/http"

	"github.com/user/shop/catalog"
)

func main() {
	http.HandleFunc("/catalog/products", productsHandler)
	http.ListenAndServe(":8080", nil)
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	p := catalog.Product{Id: 11, Name: "Taschen"}
	products := []catalog.Product{p}
	json, _ := json.Marshal(products)
	w.Write(json)
}
