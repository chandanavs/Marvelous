package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

var hero = []heros{}
var ch = make([]character, 15)
var flag int = 0

var (
	timeSumsMu sync.RWMutex
)

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
	go runDataLoop()
	router.HandleFunc("/marvels/addavengers", addHero).Methods("POST")
	router.HandleFunc("/marvels/addmutants", addHero).Methods("POST")
	router.HandleFunc("/marvels/addantiheroes", addHero).Methods("POST")
	router.HandleFunc("/marvels/{name}", getPower).Methods("GET")
	//router.HandleFunc("/marvels/mutals/{name}", getMutantPower).Methods("GET")
	//router.HandleFunc("/marvels/antihero/{name}", getAntiPower).Methods("GET")
	router.HandleFunc("/marvels", getAllCharacters).Methods("GET")
	//start web server
	err := http.ListenAndServe(":5000", router)
	if err != nil {
		fmt.Println(err)
	}

}
func runDataLoop() {
	for {
		timeSumsMu.Lock()
		for i := range hero {
			for j := range hero[i].Character {
				hero[i].Character[j].Maxpower += 5
			}
		}
		timeSumsMu.Unlock()
		time.Sleep(10 * time.Second)
	}
}

func getAllCharacters(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hero)
}
func addHero(w http.ResponseWriter, r *http.Request) {
	var newAvenger heros
	var cha character
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&newAvenger)
	for i := range newAvenger.Character {
		cha.Name = newAvenger.Character[i].Name
		cha.Maxpower = newAvenger.Character[i].Maxpower
		ch = append(ch, cha)
	}

	for i := range hero {
		if newAvenger.Name == hero[i].Name {
			for j := range newAvenger.Character {
				cha.Name = newAvenger.Character[j].Name
				cha.Maxpower = newAvenger.Character[j].Maxpower
				hero[i].Character = append(hero[i].Character, cha)
				flag = 1
			}
		}

	}
	if flag == 1 {
		json.NewEncoder(w).Encode(hero)
	} else {
		hero = append(hero, newAvenger)
		json.NewEncoder(w).Encode(hero)
	}

}

func getPower(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	for _, item := range hero {
		for _, name := range item.Character {
			if name.Name == params["name"] {
				json.NewEncoder(w).Encode(name)

			}
		}

	}
}
