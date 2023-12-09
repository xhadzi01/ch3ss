package main

import (
	"fmt"
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

// LogHandlerFuncWrapper is a simple wrapper that is able to log every function call
type LogHandlerFuncWrapper struct {
	handler http.Handler
	name    string
}

// ServeHTTP is a LogHandlerFuncWrapper overrie mtod that logs every call
func (h LogHandlerFuncWrapper) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("Func %v called\n", h.name)
	h.handler.ServeHTTP(writer, request)
}

// NewLogHandlerFuncWrapper is a constructor method used to create LogHandlerFuncWrapper
func NewLogHandlerFuncWrapper(handler http.Handler, name string) LogHandlerFuncWrapper {
	return LogHandlerFuncWrapper{
		handler: handler,
		name:    name,
	}
}

// NewURLRouter creates a new router for specified routes
func NewURLRouter(verbose bool, urlRouter URLRouter) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, urlRoute := range urlRouter.URLRoutes() {
		var handler http.Handler
		handler = urlRoute.HandlerFunctor
		if verbose {
			handler = NewLogHandlerFuncWrapper(handler, urlRoute.Name)
		}

		router.Methods(urlRoute.MethodType).
			Name(urlRoute.Name).
			Path(urlRoute.Path).
			Handler(handler)
	}

	return router
}
