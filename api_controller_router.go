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
			"MainPage",
			http.MethodGet,
			"/",
			apiRouter.ShowMainScreen,
		},
		{
			"StartNewGame",
			http.MethodPost,
			"/start-new-game",
			apiRouter.StartNewGame,
		},
		{
			"JoinGame",
			http.MethodPost,
			"/join-game",
			apiRouter.JoinGame,
		},
		{
			"WaitForOponent",
			http.MethodGet,
			"/wait-for-oponent",
			apiRouter.WaitForOponent,
		},
		{
			"ProceedToGame",
			http.MethodGet,
			"/is-ready-to-proceed",
			apiRouter.IsReadyToProceed,
		},
		{
			"ProceedToGame",
			http.MethodGet,
			"/proceed-to-game",
			apiRouter.ProceedToGame,
		},
		{
			"GetGameInfo",
			http.MethodGet,
			"/get-game-info/{SessionID}",
			apiRouter.GetGameInfo,
		},
		{
			"MoveFigure",
			http.MethodPut,
			"/move_figure/{SessionID}",
			apiRouter.MoveFigure,
		},
		{
			"GetPlayerInfo",
			http.MethodGet,
			"/get-layer-info/{PlayerID}",
			apiRouter.GetPlayerInfo,
		},
		// score overview
		{
			"GetCurrentScore",
			http.MethodGet,
			"/my-score",
			apiRouter.GetCurrentScore,
		},
		{
			"ResetScore",
			http.MethodPost,
			"/reset-score",
			apiRouter.ResetScore,
		},
		{
			"GetLeaderboard",
			http.MethodGet,
			"/leaderboard",
			apiRouter.GetLeaderboard,
		},
		// Debug functionality
		{
			"GetActiveSessions",
			http.MethodGet,
			"/debug/get-active-sessions",
			apiRouter.GetActiveSessions,
		},
		// static
		{
			"Static",
			http.MethodGet,
			"/static/{ResourceName}",
			apiRouter.GetStatic,
		},
	}
}
