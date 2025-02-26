package main

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		handler := Logger(route.HandlerFunc, route.Name)

		router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
