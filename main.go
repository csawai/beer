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

func getBeers(w http.ResponseWriter, r *http.Request) {

}
func getBeer(w http.ResponseWriter, r *http.Request) {

}

func createBeer(w http.ResponseWriter, r *http.Request) {

}

func updateBeer(w http.ResponseWriter, r *http.Request) {

}

func deleteBeer(w http.ResponseWriter, r *http.Request) {

}

func main () {

	router := mux.NewRouter()

	//handlers
	router.HandleFunc("/beer", getBeers).Methods("GET")
	router.HandleFunc("/beer/{id}", getBeer).Methods("GET")
	router.HandleFunc("/beer", createBeer).Methods("POST")
	router.HandleFunc("beer/{id}", updateBeer).Methods("POST")
	router.HandleFunc("beer/{id}", deleteBeer).Methods("DELETE")

	http.ListenAndServe(":4003", router)
	log.Fatal(http.ListenAndServe(":4003", router))



}