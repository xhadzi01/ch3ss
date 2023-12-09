package main

import (
	"fmt"
	"net/http"
)

type Controller struct {
}

func NewController() IController {
	return &Controller{}
}

func (controller *Controller) StartNewGame(http.ResponseWriter, *http.Request) {
	if controller == nil {
		panic("Controller instance is nil")
	}
	fmt.Println("StartNewGame")
}

func (controller *Controller) ProceedToGame(http.ResponseWriter, *http.Request) {
	if controller == nil {
		panic("Controller instance is nil")
	}
	fmt.Println("ProceedToGame")
}

func (controller *Controller) JoinGame(http.ResponseWriter, *http.Request) {
	if controller == nil {
		panic("Controller instance is nil")
	}
	fmt.Println("JoinGame")
}

func (controller *Controller) GetGameInfo(http.ResponseWriter, *http.Request) {
	if controller == nil {
		panic("Controller instance is nil")
	}
	fmt.Println("GetGameInfo")
}

func (controller *Controller) GetPlayerInfo(http.ResponseWriter, *http.Request) {
	if controller == nil {
		panic("Controller instance is nil")
	}
	fmt.Println("GetPlayerInfo")
}
