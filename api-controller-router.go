package main

import (
	"net/http"
)

type APIControllerRouter struct {
	IController
}

func NewAPIControllerRouter(controller IController) *APIControllerRouter {
	if controller == nil {
		panic("Controller is nil")
	}
	return &APIControllerRouter{
		IController: controller,
	}
}

func (apiRouter *APIControllerRouter) URLRoutes() URLRoutes {
	if apiRouter == nil || apiRouter.IController == nil {
		panic("Controller is nil")
	}

	return URLRoutes{
		{
			"StartNewGame",
			http.MethodPost,
			"/start_new_game",
			apiRouter.StartNewGame,
		},
		{
			"ProceedToGame",
			http.MethodPost,
			"/proceed_to_game/{SessionID}",
			apiRouter.ProceedToGame,
		},
		{
			"JoinGame",
			http.MethodPost,
			"/join_game/{SessionID}",
			apiRouter.JoinGame,
		},
		{
			"GetGameInfo",
			http.MethodGet,
			"/get_game_info/{SessionID}",
			apiRouter.GetGameInfo,
		},
		{
			"GetPlayerInfo",
			http.MethodGet,
			"/get_layer_info/{PlayerID}",
			apiRouter.GetPlayerInfo,
		},
	}
}
