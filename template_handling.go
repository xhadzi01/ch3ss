package main

import (
	"errors"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

// LoadTemplates is a helper funclion that allows loading static templates
func LoadTemplates(temmplateNames []string) (templ *template.Template, err error) {
	if len(temmplateNames) == 0 {
		err = errors.New("could not create template loader, no template files specified")
		return
	}

	if curDir, errCwd := os.Getwd(); errCwd != nil {
		err = fmt.Errorf("coud not retrieve current working directory, reason: %v", errCwd)
	} else {
		retvalTempl := template.New("")

		for _, filename := range temmplateNames {
			fullFilePath := filepath.Join(curDir, "static", "template", filename)
			fileInfo, statErr := os.Stat(fullFilePath)
			if os.IsNotExist(statErr) || fileInfo.IsDir() {
				err = fmt.Errorf("coud not retrieve template file: %v", fullFilePath)
				return
			}

			if _, errParse := retvalTempl.ParseFiles(fullFilePath); errParse != nil {
				err = fmt.Errorf("coud not parse template file, reason: %v", errParse)
				return
			}
		}

		templ = retvalTempl
	}
	return
}
