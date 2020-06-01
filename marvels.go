package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var avenger []heros = []heros{}
var mutant []heros = []heros{}
var antiheroes []heros = []heros{}

type character struct {
	Name     string `json:"name"`
	Maxpower int    `json:"max_power"`
}
type heros struct {
	Name      string      `json:"name"`
	Character []character `json:"character"`
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/marvels/addavengers", addAvenger).Methods("POST")
	router.HandleFunc("/marvels/addmutants", addMutants).Methods("POST")
	router.HandleFunc("/marvels/addantiheroes", addAntiHeroes).Methods("POST")
	router.HandleFunc("/marvels", getAllCharacters).Methods("GET")
	//start web server
	err := http.ListenAndServe(":5000", router)
	if err != nil {
		fmt.Println(err)
	}

}
func getAllCharacters(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(avenger)
	json.NewEncoder(w).Encode(mutant)
	json.NewEncoder(w).Encode(antiheroes)
}
func addAvenger(w http.ResponseWriter, r *http.Request) {
	var newAvenger heros
	json.NewDecoder(r.Body).Decode(&newAvenger)
	avenger = append(avenger, newAvenger)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(avenger)
}

func addMutants(w http.ResponseWriter, r *http.Request) {
	var newMutant heros
	json.NewDecoder(r.Body).Decode(&newMutant)
	mutant = append(mutant, newMutant)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mutant)
}

func addAntiHeroes(w http.ResponseWriter, r *http.Request) {
	var newAnti heros
	json.NewDecoder(r.Body).Decode(&newAnti)
	antiheroes = append(antiheroes, newAnti)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newAnti)
}
