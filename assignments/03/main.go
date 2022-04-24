package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"time"
)

type Product struct {
	ID     string `json:"ID"`
	name   string `json:"name"`
	price  string `json:"price"`
	amount string `json:"amount"`
}

var Products []Product

/*
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}
*/

func returnAllProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllProducts")
	json.NewEncoder(w).Encode(Products)
}

func getProduct() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(r)
		ID := params["ID"]
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(ID)

		err := productCollection.FindOne(ctx, bson.M{"id": ID}).Decode(&ID)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		var result bson.M
		fmt.Println(result)
	}
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	//router.HandleFunc("/", homePage)
	router.HandleFunc("/products/list", returnAllProducts).Methods("LIST")
	router.HandleFunc("/product/{Name}", getProduct).Methods("GET")
	router.HandleFunc("/product/{Name}", updateProduct).Methods("UPDATE")
	router.HandleFunc("/product/{Name}", deleteProduct).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}

//app.Post("/product", CreateProduct)

func main() {
	client := ConnectDB()
	GetCollection(client, "products")
	handleRequests()
	/*
		Products = []Product{
			{name: "Ford Mustang", price: "$20000", amount: "5"},
			{name: "Chevrolet Camaro", price: "$30000", amount: "8"},
		}
	*/
}
