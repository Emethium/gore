package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Pokemon struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Typeone string `json:"typeone,omitempty"`
	Typetwo string `json:"typetwo,omitempty"`
}

var pokemanz []Pokemon

func GetPokemonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range pokemanz {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Pokemon{})
}

func GetPokemanzEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(pokemanz)
}

func CreatePokemonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var pokemon Pokemon
	_ = json.NewDecoder(req.Body).Decode(&pokemon)
	pokemon.ID = params["id"]
	pokemanz = append(pokemanz, pokemon)
	json.NewEncoder(w).Encode(pokemanz)
}

func DeletePokemonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range pokemanz {
		if item.ID == params["id"] {
			pokemanz = append(pokemanz[:index], pokemanz[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(pokemanz)
}

func main() {
	router := mux.NewRouter()

	// Populating pokemon
	pokemanz = append(pokemanz, Pokemon{ID: "1", Name: "Bulbasaur", Typeone: "Grass", Typetwo: "Poison"})
	pokemanz = append(pokemanz, Pokemon{ID: "155", Name: "Cyndaquil", Typeone: "Fire", Typetwo: ""})
	pokemanz = append(pokemanz, Pokemon{ID: "25", Name: "Pikachu", Typeone: "Electric", Typetwo: ""})

	// Defining pokemon routes
	router.HandleFunc("/pokemon", GetPokemanzEndpoint).Methods("GET")
	router.HandleFunc("/pokemon/{id}", GetPokemonEndpoint).Methods("GET")
	router.HandleFunc("/pokemon/{id}", CreatePokemonEndpoint).Methods("POST")
	router.HandleFunc("/pokemon/{id}", DeletePokemonEndpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":12345", router))
}
