package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Superhero struct {
	Id          int    `json:"id"`
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
	w.Header().Set("Content-Type", "application/json")
	var hero Superhero
	_ = json.NewDecoder(r.Body).Decode(&hero)
	hero.Id = rand.Intn(1000000)
	superheroes = append(superheroes, hero)
	json.NewEncoder(w).Encode(&hero)

}

func deleteSuperHeroes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, hero := range superheroes {
		if strconv.Itoa(hero.Id) == params["id"] {
			superheroes = append(superheroes[:index], superheroes[index+1:]...)
			json.NewEncoder(w).Encode(hero)
			break
		}
	}
	json.NewEncoder(w).Encode("deletion successful")
}

func getSuperHeroeByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, hero := range superheroes {
		if strconv.Itoa(hero.Id) == params["id"] {
			json.NewEncoder(w).Encode(hero)
			return
		}
	}
	json.NewEncoder(w).Encode(&Superhero{})
}

func greet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Welcome to the superhero page!!")
}

func main() {
	superheroes = append(superheroes, Superhero{1, "Ironman", "Ironsuit", "USA", 45})
	superheroes = append(superheroes, Superhero{2, "Hulk", "super-strength", "USA", 43})

	r := mux.NewRouter()
	r.HandleFunc("/superhero/greet", greet).Methods("GET")
	r.HandleFunc("/superhero/all", getAllSuperHeroes).Methods("GET")

	r.HandleFunc("/superhero/{id}", getSuperHeroeByName).Methods("GET")
	r.HandleFunc("/superhero/create", createSuperHeroes).Methods("POST")
	r.HandleFunc("/superhero/delete/{id}", deleteSuperHeroes).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))

}
