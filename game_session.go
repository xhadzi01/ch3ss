package main

type GameSession struct {
}

func NewGameSession() IGameSession {
	return &GameSession{}
}
