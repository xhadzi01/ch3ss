package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Controller struct {
	Sessions
}

func NewController() IController {
	return &Controller{
		Sessions: make(Sessions, 0),
	}
}

func (controller *Controller) StartNewGame(writter http.ResponseWriter, request *http.Request) {
	if controller == nil {
		panic("Controller instance is nil")
	}
	session := NewSession()
	controller.Sessions = append(controller.Sessions, session)

	encodeResponseAsJSON(writter, http.StatusOK, struct {
		SessionID
		PlayerID
	}{
		SessionID: session.SessionID,
		PlayerID:  session.Player1ID,
	})
}

func (controller *Controller) ProceedToGame(writter http.ResponseWriter, request *http.Request) {
	if controller == nil {
		panic("Controller instance is nil")
	}
	fmt.Println("ProceedToGame")
}

func (controller *Controller) JoinGame(writter http.ResponseWriter, request *http.Request) {
	if controller == nil {
		panic("Controller instance is nil")
	}
	fmt.Println("JoinGame")
}

func (controller *Controller) GetGameInfo(writter http.ResponseWriter, request *http.Request) {
	if controller == nil {
		panic("Controller instance is nil")
	}
	fmt.Println("GetGameInfo")
}

func (controller *Controller) GetPlayerInfo(writter http.ResponseWriter, request *http.Request) {
	if controller == nil {
		panic("Controller instance is nil")
	}
	fmt.Println("GetPlayerInfo")
}

func (controller *Controller) GetActiveSessions(writter http.ResponseWriter, request *http.Request) {
	if controller == nil {
		panic("Controller instance is nil")
	}
	fmt.Println("GetPlayerInfo")
}

func encodeResponseAsJSON(w http.ResponseWriter, statusCode int, itf interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(itf)
}
