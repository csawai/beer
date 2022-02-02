package main

import (
	
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Beer struct {

	ID string `json:"id"`
	ImageNumber string `json:"ImageNumber"`
	Name string `json:"Name"`
	BeerType string `json:"BeerType"`
}

var beers []Beer

func main () {

	router := mux.NewRouter()

	//handlers
	router.HandleFunc("/beer", getBeers).Methods("GET")
	router.HandleFunc("/beer/{id}", getBeer).Methods("GET")
	router.HandleFunc("/beer", createBeer).Methods("POST")
	router.HandleFunc("beer/{id}", updateBeer).Methods("POST")
	router.HandleFunc("beer/{id}", deleteBeer).Methods("POST")



}