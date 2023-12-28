package main

var player1defaultFigurePositions PlayerFigurePositions
var player2defaultFigurePositions PlayerFigurePositions

func init() {
	player1defaultFigurePositions = PlayerFigurePositions{
		pawn1:   9,
		pawn2:   10,
		pawn3:   11,
		pawn4:   12,
		pawn5:   13,
		pawn6:   14,
		pawn7:   15,
		pawn8:   16,
		rook1:   1,
		knight1: 2,
		bishop1: 3,
		king:    4,
		queen:   5,
		bishop2: 6,
		knight2: 7,
		rook2:   8,
	}
	player2defaultFigurePositions = PlayerFigurePositions{
		pawn1:   56,
		pawn2:   55,
		pawn3:   54,
		pawn4:   53,
		pawn5:   52,
		pawn6:   51,
		pawn7:   50,
		pawn8:   49,
		rook1:   64,
		knight1: 63,
		bishop1: 62,
		king:    61,
		queen:   60,
		bishop2: 59,
		knight2: 58,
		rook2:   57,
	}
}
