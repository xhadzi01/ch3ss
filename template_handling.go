package main

import (
	"errors"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

// Page is a helper that allows setting of some special settings related to http page
type Page struct {
	Title string
}

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
			fullFilePath := filepath.Join(curDir, "static", "templates", filename)
			fileInfo, statErr := os.Stat(fullFilePath)
			if os.IsNotExist(statErr) || !fileInfo.IsDir() {
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
