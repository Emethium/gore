package main

type Pokemon struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Typeone string `json:"typeone"`
	Typetwo string `json:"typetwo"`
}

type Pokemanz []Pokemon
