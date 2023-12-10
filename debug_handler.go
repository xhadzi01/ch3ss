package main

import (
	"fmt"
	"net/http"
)

type DebugHandler struct {
	handler http.Handler
}

func NewDebugHandler(handler http.Handler) http.Handler {
	return &DebugHandler{handler: handler}
}

func (dh *DebugHandler) ServeHTTP(writter http.ResponseWriter, request *http.Request) {
	fmt.Println("--------------------------------------")
	fmt.Printf("\tMethod=%v\n", request.Method)
	fmt.Printf("\tURL=%v\n", request.URL)
	fmt.Printf("\tHeaders=%v\n", request.Header)
	fmt.Printf("\tContentLength=%v\n", request.ContentLength)
	fmt.Println("--------------------------------------")
	dh.handler.ServeHTTP(writter, request)
}
