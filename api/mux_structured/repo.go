package main

import "fmt"

var currentId int

var pokemanz Pokemanz

// Give us some seed data
func init() {
	RepoCreate(Pokemon{Id: 1, Name: "Bulbasaur", Typeone: "Grass", Typetwo: "Poison"})
	RepoCreate(Pokemon{Id: 2, Name: "Ivysaur", Typeone: "Grass", Typetwo: "Poison"})
	RepoCreate(Pokemon{Id: 3, Name: "Venusaur", Typeone: "Grass", Typetwo: "Poison"})
}

func RepoFind(id int) Pokemon {
	for _, p := range pokemanz {
		if p.Id == id {
			return p
		}
	}
	// return empty Todo if not found
	return Pokemon{}
}

func RepoCreate(p Pokemon) Pokemon {
	currentId++
	p.Id = currentId
	pokemanz = append(pokemanz, p)
	return p
}

func RepoDestroy(id int) error {
	for i, p := range pokemanz {
		if p.Id == id {
			pokemanz = append(pokemanz[:i], pokemanz[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Pokemon with id of %d to delete", id)
}
