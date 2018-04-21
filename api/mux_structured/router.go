package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	// StrictSlash defines the trailing slash behavior for new routes. The initial
	// value is false.
	//
	// When true, if the route path is "/path/", accessing "/path" will perform a redirect
	// to the former and vice versa.
	//
	// When false, if the route path is "/path", accessing "/path/" will not match
	// this route and vice versa.
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}
