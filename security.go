package main

import (
	"os"
	"path/filepath"
)

var (
	DefaultCertFilePath string
	DefaultKeyFilePath  string
)

func init() {
	executableLocation, err := os.Executable()
	if err != nil {
		panic(err)
	}
	executableDirectory := filepath.Dir(executableLocation)
	DefaultCertFilePath = filepath.Join(executableDirectory, "server-cert.pem")
	DefaultKeyFilePath = filepath.Join(executableDirectory, "server-key.pem")
}
