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
	Name   string `bson:"Name" json:"Name"`
	Value  int32  `bson:"Value" json:"Value"`
	Amount int32  `bson:"Amount" json:"Amount"`
}

var Products []Product

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CTC-API cars")
}

func returnAllProducts(w http.ResponseWriter, r *http.Request) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.Find(ctx, bson.M{})
	result.All(ctx, &Products)
	json.NewEncoder(w).Encode(Products)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	params := mux.Vars(r)
	carName := params["Name"]
	var car bson.M
	collection.FindOne(ctx, bson.M{"Name": carName}).Decode(&car)
	json.NewEncoder(w).Encode(car)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	//TODO
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	params := mux.Vars(r)
	carName := params["Name"]
	result, _ := collection.DeleteOne(ctx, bson.M{"Name": carName})
	json.NewEncoder(w).Encode(result)
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	//router.HandleFunc("/home", homePage).Methods("GET")
	router.HandleFunc("/products/list", returnAllProducts).Methods("GET")
	router.HandleFunc("/product/{Name}", getProduct).Methods("GET")
	router.HandleFunc("/product/{Name}", updateProduct).Methods("UPDATE")
	router.HandleFunc("/product/{Name}", deleteProduct).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func routerSetup() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://ahoj:ahoj@cluster0.pfaom.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	collection = client.Database("ctc").Collection("cars")
}

var collection *mongo.Collection

func main() {
	routerSetup()
	handleRequests()
}
