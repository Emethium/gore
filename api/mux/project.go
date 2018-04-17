package main

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

type Person struct {
    ID        string   `json:"id,omitempty"`
    Firstname string   `json:"firstname,omitempty"`
    Lastname  string   `json:"lastname,omitempty"`
    Address   *Address `json:"address,omitempty"`
}

type Address struct {
    City  string `json:"city,omitempty"`
    State string `json:"state,omitempty"`
}

type Pokemon struct {
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Typeone string `json:"typeone,omitempty"`
	Typetwo string `json:"typetwo,omitempty"`
}

var people []Person
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

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for _, item := range people {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Person{})
}

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
    json.NewEncoder(w).Encode(people)
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    var person Person
    _ = json.NewDecoder(req.Body).Decode(&person)
    person.ID = params["id"]
    people = append(people, person)
    json.NewEncoder(w).Encode(people)
}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for index, item := range people {
        if item.ID == params["id"] {
            people = append(people[:index], people[index+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(people)
}

func main() {
    router := mux.NewRouter()
	
	// Populating examples
	people = append(people, Person{ID: "1", Firstname: "Nic", Lastname: "Raboy", Address: &Address{City: "Dublin", State: "CA"}})
	people = append(people, Person{ID: "2", Firstname: "Maria", Lastname: "Raboy"})
	pokemanz = append(pokemanz, Pokemon{ID: "1", Name: "Bulbasaur", Typeone: "Grass", Typetwo: "Poison"})
	pokemanz = append(pokemanz, Pokemon{ID: "155", Name: "Cyndaquil", Typeone: "Fire", Typetwo: ""})
	pokemanz = append(pokemanz, Pokemon{ID: "25", Name: "Pikachu", Typeone: "Electric", Typetwo: ""})
	
	// Defining people routes 
    router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
    router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
    router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
	
	// Defining pokemon routes
	router.HandleFunc("/pokemon", GetPokemanzEndpoint).Methods("GET")
    router.HandleFunc("/pokemon/{id}", GetPokemonEndpoint).Methods("GET")
    router.HandleFunc("/pokemon/{id}", CreatePokemonEndpoint).Methods("POST")
	router.HandleFunc("/pokemon/{id}", DeletePokemonEndpoint).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":12345", router))
}