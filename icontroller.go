package main

import "net/http"

type IController interface {
	StartNewGame(http.ResponseWriter, *http.Request)
	ProceedToGame(http.ResponseWriter, *http.Request)
	JoinGame(http.ResponseWriter, *http.Request)
	GetGameInfo(http.ResponseWriter, *http.Request)
	GetPlayerInfo(http.ResponseWriter, *http.Request)
	// debug
	GetActiveSessions(http.ResponseWriter, *http.Request)
}
