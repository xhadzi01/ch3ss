package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
)

func httpRequestHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello,World!\n"))
}

func Listen() {
	server := http.Server{
		Addr:    ":5555",
		Handler: http.HandlerFunc(httpRequestHandler),
	}

	defer server.Close()
	log.Fatal(server.ListenAndServe())
}

func ListenTLS() {
	// load tls certificates
	serverTLSCert, err := tls.LoadX509KeyPair(CertFilePath, KeyFilePath)
	if err != nil {
		panic(fmt.Sprintf("Error loading certificate and key file: %v", err))
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{serverTLSCert},
	}

	server := http.Server{
		Addr:      ":5555",
		Handler:   http.HandlerFunc(httpRequestHandler),
		TLSConfig: tlsConfig,
	}

	defer server.Close()
	log.Fatal(server.ListenAndServeTLS("", ""))
}

func main() {
	args := os.Args
	if (len(args) > 1) && (args[1] == "TLS") {
		ListenTLS()
	} else {
		Listen()
	}
}
