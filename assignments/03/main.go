package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type Product struct {
	name   string `json:"name"`
	price  string `json:"price"`
	amount string `json:"amount"`
}

var Products []Product

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func returnAllProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllProducts")
	json.NewEncoder(w).Encode(Products)
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	ConnectDB()
	router.HandleFunc("/", homePage)
	router.HandleFunc("/all", returnAllProducts)
	log.Fatal(http.ListenAndServe(":8080", router))
}

//app.Post("/product", CreateProduct)

func main() {
	handleRequests()
	fmt.Println("Rest API v2.0 - Mux Routers")
	Products = []Product{
		Product{name: "Ford Mustang", price: "$20000", amount: "5"},
		Product{name: "Chevrolet Camaro", price: "$30000", amount: "8"},
	}

	dbURL := os.ExpandEnv("mongodb+srv://admin:<password>@cluster0.ymdy2.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
	fmt.Println("DB URL: ", dbURL)

}
