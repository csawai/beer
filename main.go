package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Beer struct {
	ID          string `json:"id"`
	ImageNumber string `json:"ImageNumber"`
	Name        string `json:"Name"`
	BeerType    string `json:"BeerType"`
}

var beers []Beer

func getBeers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(beers)

	
}
func getBeer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for _, item := range beers {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}

}

func createBeer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var newBeer Beer
	json.NewDecoder(r.Body).Decode(&newBeer)
	newBeer.ID = strconv.Itoa(len(beers) + 1)
	beers = append(beers, newBeer)
	json.NewEncoder(w).Encode(newBeer)

}

func updateBeer(w http.ResponseWriter, r *http.Request) {

}

func deleteBeer(w http.ResponseWriter, r *http.Request) {

}

func main() {
	beers = append(beers, Beer{ID: "1", ImageNumber: "hello", Name: " beer", BeerType: "ALe"})
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
