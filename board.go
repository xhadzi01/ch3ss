package main

type Board struct {
	player1 *PlayerFigurePositions
	player2 *PlayerFigurePositions
}

func NewBoard() *Board {
	return &Board{
		player1: NewPlayerFigurePositions(Player1),
		player2: NewPlayerFigurePositions(Player2),
	}
}
