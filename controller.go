package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Controller struct {
	behavior IBehaviorModel
}

func NewController(bhv IBehaviorModel) IController {
	if bhv == nil {
		panic("could not create controller, behavior is invalid(nil)")
	}
	return &Controller{
		behavior: bhv,
	}
}

func (controller *Controller) StartNewGame(writter http.ResponseWriter, request *http.Request) {
	if controller == nil {
		panic("Controller instance is nil")
	}

	// update logic
	session, startGameErr := controller.behavior.StartNewGame()
	if startGameErr != nil {
		encodeResponseAsJSON(writter, http.StatusBadRequest, startGameErr)
		return
	}

	// respond
	encodeResponseAsJSON(writter, http.StatusOK, startNewGameResponseModel{
		SessionID:    uint64(session.SessionID),
		SessionToken: string(session.SessionToken),
		PlayerID:     uint64(session.Player1ID),
	})
}

func (controller *Controller) JoinGame(writter http.ResponseWriter, request *http.Request) {
	if controller == nil {
		panic("Controller instance is nil")
	}

	// parse input
	var joinGameRequest joinGameRequestModel
	if errParsing := parseJSONMessage(&joinGameRequest, request); errParsing != nil {
		encodeResponseAsJSON(writter, http.StatusBadRequest, errParsing)
		return
	}

	// update logic
	session, joinGameErr := controller.behavior.JoinGame(SessionID(joinGameRequest.SessionID))

	// respond
	var responseModel joinGameResponseModel

	if joinGameErr == nil {
		responseModel = joinGameResponseModel{
			PlayerID:     uint64(*session.Player2ID),
			SessionToken: string(session.SessionToken),
		}
	} else {
		responseModel = joinGameResponseModel{
			FailedReason: joinGameErr.Error(),
		}
	}

	encodeResponseAsJSON(writter, http.StatusOK, responseModel)
}

func (controller *Controller) ProceedToGame(writter http.ResponseWriter, request *http.Request) {
	if controller == nil {
		panic("Controller instance is nil")
	}

	// parse input
	var proceedToGameRequest proceedToGameGameRequestModel
	if errParsing := parseJSONMessage(&proceedToGameRequest, request); errParsing != nil {
		encodeResponseAsJSON(writter, http.StatusBadRequest, errParsing)
		return
	}

	// update logic
	proceed, poceedToGameErr := controller.behavior.ProceedToGame(SessionID(proceedToGameRequest.SessionID), SessionToken(proceedToGameRequest.SessionToken), PlayerID(proceedToGameRequest.PlayerID))

	// respond
	if poceedToGameErr == nil {
		encodeResponseAsJSON(writter, http.StatusOK, proceedToGameGameResponseModel{
			Proceed: proceed,
		})
	} else {
		encodeResponseAsJSON(writter, http.StatusBadRequest, poceedToGameErr.Error())
	}
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
