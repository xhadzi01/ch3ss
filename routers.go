package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// URLRoute defines the name, method, pattern and handler for routing
type URLRoute struct {
	Name           string
	MethodType     string
	Path           string
	HandlerFunctor http.HandlerFunc
}

// URLRoutes is a collection of all available URL routes
type URLRoutes []URLRoute

// URLRouter defines the required methods for retrieving api routes
type URLRouter interface {
	URLRoutes() URLRoutes
}

// NewURLRouter creates a new router for specified routes
func NewURLRouter(urlRouter URLRouter) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, urlRoute := range urlRouter.URLRoutes() {

		router.Methods(urlRoute.MethodType).
			Name(urlRoute.Name).
			Path(urlRoute.Path).
			Handler(urlRoute.HandlerFunctor)
	}

	return router
}
