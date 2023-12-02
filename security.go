package main

import (
	"os"
	"path/filepath"
)

var (
	CertFilePath string
	KeyFilePath  string
)

func init() {
	executableLocation, err := os.Executable()
	if err != nil {
		panic(err)
	}
	executableDirectory := filepath.Dir(executableLocation)
	CertFilePath = filepath.Join(executableDirectory, "server-cert.pem")
	KeyFilePath = filepath.Join(executableDirectory, "server-key.pem")
}
