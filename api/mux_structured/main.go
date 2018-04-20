package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// StrictSlash defines the trailing slash behavior for new routes. The initial
	// value is false.
	//
	// When true, if the route path is "/path/", accessing "/path" will perform a redirect
	// to the former and vice versa.
	//
	// When false, if the route path is "/path", accessing "/path/" will not match
	// this route and vice versa.
	router := mux.NewRouter().StrictSlash(true)

	log.Fatal(http.ListenAndServe(":8080", router))
}
