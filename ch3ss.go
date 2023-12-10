package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"go.elastic.co/apm/module/apmgorilla"
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
	if config.TLSConnection {
		fmt.Printf("\tcert-path: '%v'\n", config.CertFilePath)
		fmt.Printf("\tkey-path: '%v'\n", config.KeyFilePath)
	}
	fmt.Printf("\tverbose: '%v'\n", config.VerboseLogging)

	return config
}

func Listen(bindingAddress string, handler http.Handler, verbose bool) {
	server := http.Server{
		Addr:    bindingAddress,
		Handler: handler,
	}

	defer server.Close()
	log.Fatal(server.ListenAndServe())
}

func ListenTLS(bindingAddress string, handler http.Handler, certFilePath string, keyFilePath string, verbose bool) {
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
		Handler:   handler,
		TLSConfig: tlsConfig,
	}

	defer server.Close()
	log.Fatal(server.ListenAndServeTLS("", ""))
}

func main() {
	fmt.Println(os.Args)
	log.Printf("Server is starting")
	config := ParseConfigurationFromCmd()

	controller := NewController()
	apiControllerRouter := NewAPIControllerRouter(controller)
	router := NewURLRouter(apiControllerRouter)
	router.Use(apmgorilla.Middleware())

	var handler http.Handler = router
	if config.VerboseLogging {
		handler = NewDebugHandler(handler)
	}

	if config.TLSConnection {
		ListenTLS(config.BindingAddress, handler, config.CertFilePath, config.KeyFilePath, config.VerboseLogging)
	} else {
		Listen(config.BindingAddress, handler, config.VerboseLogging)
	}
}
