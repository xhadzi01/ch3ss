package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func parseJSONMessage(retval interface{}, request *http.Request) (err error) {
	if request == nil {
		err = errors.New("input reqest is invalid(nil)")
		return
	} else if request.Body == nil {
		err = errors.New("request contains no body")
		return
	}

	defer request.Body.Close()

	if bodyText, readAllErr := io.ReadAll(request.Body); readAllErr != nil {
		err = fmt.Errorf("could not read body content. Reason: %v", readAllErr)
	} else if err = json.Unmarshal(bodyText, &retval); err != nil {
		err = fmt.Errorf("could not read body content. Reason: %v", readAllErr)
	}
	return
}
