package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
)

type Configuration struct {
	BindingAddress string
	TLSConnection  bool
	CertFilePath   string
	KeyFilePath    string
	VerboseLogging bool
}

func NewConfiguration() *Configuration {

	return &Configuration{
		BindingAddress: ":5555",
		TLSConnection:  false,
		CertFilePath:   DefaultCertFilePath,
		KeyFilePath:    DefaultKeyFilePath,
		VerboseLogging: false,
	}
}

func ParseConfigurationFromCmd() *Configuration {
	config := NewConfiguration()

	flag.StringVar(&config.BindingAddress, "binding-address", config.BindingAddress, "address on which this service should listen")
	flag.BoolVar(&config.TLSConnection, "use-tls", config.TLSConnection, "enable TLS connection")
	flag.StringVar(&config.CertFilePath, "cert-path", config.CertFilePath, "location of cert file, used only if `use-tls` is set")
	flag.StringVar(&config.KeyFilePath, "key-path", config.KeyFilePath, "location of key file, used only if `use-tls` is set")
	flag.BoolVar(&config.VerboseLogging, "verbose", config.VerboseLogging, "enable verbose logging")
	flag.Parse()

	fmt.Println("Configuration:")
	fmt.Printf("\tbinding-address:'%v'", config.BindingAddress)
	fmt.Printf("\tuse-tls: '%v'\n", config.TLSConnection)
	fmt.Printf("\tcert-path: '%v'\n", config.CertFilePath)
	fmt.Printf("\tkey-path: '%v'\n", config.KeyFilePath)
	fmt.Printf("\tverbose: '%v'\n", config.VerboseLogging)

	return config
}

func httpRequestHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello,World!\n"))
}

func Listen(bindingAddress string, verbose bool) {
	server := http.Server{
		Addr:    bindingAddress,
		Handler: http.HandlerFunc(httpRequestHandler),
	}

	defer server.Close()
	log.Fatal(server.ListenAndServe())
}

func ListenTLS(bindingAddress, certFilePath, keyFilePath string, verbose bool) {
	// load tls certificates
	serverTLSCert, err := tls.LoadX509KeyPair(certFilePath, keyFilePath)
	if err != nil {
		panic(fmt.Sprintf("Error loading certificate and key file: %v", err))
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{serverTLSCert},
	}

	server := http.Server{
		Addr:      bindingAddress,
		Handler:   http.HandlerFunc(httpRequestHandler),
		TLSConfig: tlsConfig,
	}

	defer server.Close()
	log.Fatal(server.ListenAndServeTLS("", ""))
}

func main() {
	log.Printf("Server is starting")
	config := ParseConfigurationFromCmd()

	if config.TLSConnection {
		ListenTLS(config.BindingAddress, config.CertFilePath, config.KeyFilePath, config.VerboseLogging)
	} else {
		Listen(config.BindingAddress, config.VerboseLogging)
	}
}
