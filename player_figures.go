package main

type PlayerFigures struct {
	pawn1   uint64
	pawn2   uint64
	pawn3   uint64
	pawn4   uint64
	pawn5   uint64
	pawn6   uint64
	pawn7   uint64
	pawn8   uint64
	rook1   uint64
	knight1 uint64
	bishop1 uint64
	queen   uint64
	king    uint64
	bishop2 uint64
	knight2 uint64
	rook2   uint64
}

type PlayerType uint8

const (
	Player1 PlayerType = 0
	Player2 PlayerType = 1
)

type PlayerFigureOutput byte

const (
	pawnFigureOutput   PlayerFigureOutput = 'P'
	rookFigureOutput   PlayerFigureOutput = 'R'
	knightFigureOutput PlayerFigureOutput = 'K'
	bishopFigureOutput PlayerFigureOutput = 'B'
	queenFigureOutput  PlayerFigureOutput = 'Q'
	kingFigureOutput   PlayerFigureOutput = 'G'
)

func NewPlayerFigures(pt PlayerType) *PlayerFigures {
	if pt == Player1 {
		retval := player1defaultFigures
		return &retval
	} else if pt == Player2 {
		retval := player2defaultFigures
		return &retval
	}

	panic("unknown player type")
}
