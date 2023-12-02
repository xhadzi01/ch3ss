package main

type Board struct {
	player1 *PlayerFigures
	player2 *PlayerFigures
}

func NewBoard() *Board {
	return &Board{
		player1: NewPlayerFigures(Player1),
		player2: NewPlayerFigures(Player2),
	}
}
