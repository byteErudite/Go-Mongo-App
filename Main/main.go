package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Superhero struct {
	Name        string `json:"name"`
	Superpower  string `json:"superpower"`
	Nationality string `json:"nationality"`
	Age         int    `json:"age"`
}

var superheroes []Superhero

func getAllSuperHeroes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(superheroes)
}

func createSuperHeroes(w http.ResponseWriter, r *http.Request) {

}

func deleteSuperHeroes(w http.ResponseWriter, r *http.Request) {

}

func getSuperHeroeByName(w http.ResponseWriter, r *http.Request) {

}

func main() {
	superheroes = append(superheroes, Superhero{"Ironman", "Ironsuit", "USA", 45})
	superheroes = append(superheroes, Superhero{"Hulk", "super-strength", "USA", 43})

	r := mux.NewRouter()
	r.HandleFunc("/superhero/all", getAllSuperHeroes).Methods("GET")
	r.HandleFunc("/superhero/all", getSuperHeroeByName).Methods("GET")
	r.HandleFunc("/superhero/create", createSuperHeroes).Methods("POST")
	r.HandleFunc("/superhero/delete/{name}", deleteSuperHeroes).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))

}
