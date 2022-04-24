package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func returnAllProducts(w http.ResponseWriter, r *http.Request) {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	result, _ := collection.Find(ctx, bson.M{})
	result.All(ctx, &Products)
	fmt.Println("Endpoint Hit: returnAllProducts")
	fmt.Println(Products)
	json.NewEncoder(w).Encode(Products)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["name"]
	fmt.Println(ID)
	fmt.Println("Endpoint Hit: returnAllProducts")
	for _, singleEvent := range Products {
		if singleEvent.ID == ID {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/products/list", returnAllProducts).Methods("GET")
	router.HandleFunc("/product/{Name}", getProduct).Methods("GET")
	//router.HandleFunc("/product/{Name}", updateProduct).Methods("UPDATE")
	//router.HandleFunc("/product/{Name}", deleteProduct).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func routerSetup() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0.ymdy2.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	collection = client.Database("ctcAPI").Collection("products")
}

//app.Post("/product", CreateProduct)
var collection *mongo.Collection

func main() {
	routerSetup()
	handleRequests()
	/*
		Products = []Product{
			{name: "Ford Mustang", price: "$20000", amount: "5"},
			{name: "Chevrolet Camaro", price: "$30000", amount: "8"},
		}
	*/

}
