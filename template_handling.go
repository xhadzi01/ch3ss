package main

import (
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
)

// LoadTemplates is a helper funclion that allows loading static templates
func LoadTemplates(temmplateNames []string) (templ *template.Template, err error) {
	if len(temmplateNames) == 0 {
		err = errors.New("could not create template loader, no template files specified")
		return
	}

	retvalTempl := template.New("")

	for _, filename := range temmplateNames {
		if fullFilePath, pathErr := getResourcePath("template", filename); pathErr != nil {
			err = pathErr
			return
		} else if _, errParse := retvalTempl.ParseFiles(fullFilePath); errParse != nil {
			err = fmt.Errorf("coud not parse template file, reason: %v", errParse)
			return
		}
	}

	templ = retvalTempl
	return
}

// helper function retrieving file content as a byte data
func LoadStatic(subpaths ...string) (data []byte, err error) {
	if fullPath, resourceErr := getResourcePath(subpaths...); resourceErr != nil {
		err = resourceErr
	} else if redData, readErr := ioutil.ReadFile(fullPath); readErr != nil {
		err = readErr
	} else {
		data = redData
	}

	return
}

func getResourcePath(subpaths ...string) (fullPath string, err error) {
	curDir, errCwd := os.Getwd()
	if errCwd != nil {
		err = fmt.Errorf("coud not retrieve current working directory, reason: %v", errCwd)
		return
	}

	fullFilePath := filepath.Join(curDir, "static")
	for _, subpath := range subpaths {
		fullFilePath = filepath.Join(fullFilePath, subpath)
	}

	fileInfo, statErr := os.Stat(fullFilePath)
	if os.IsNotExist(statErr) || fileInfo.IsDir() {
		err = fmt.Errorf("coud not find file: %v", fullFilePath)
		return
	}

	fullPath = fullFilePath
	return
}
