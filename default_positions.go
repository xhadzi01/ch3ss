package main

var player1defaultFigurePositions PlayerFigurePositions
var player2defaultFigurePositions PlayerFigurePositions

func init() {
	player1defaultFigurePositions = PlayerFigurePositions{
		[]FigureInfo{
			{ // pawn1
				PawnFigure,
				9,
				9,
			},
			{ // pawn2
				PawnFigure,
				10,
				10,
			},
			{ // pawn3
				PawnFigure,
				11,
				11,
			},
			{ // pawn4
				PawnFigure,
				12,
				12,
			},
			{ // pawn5
				PawnFigure,
				13,
				13,
			},
			{ // pawn6
				PawnFigure,
				14,
				14,
			},
			{ // pawn7
				PawnFigure,
				15,
				15,
			},
			{ // pawn8
				PawnFigure,
				16,
				16,
			},
			{ // rook1
				RookFigure,
				1,
				1,
			},
			{ // knight1
				KnightFigure,
				2,
				2,
			},
			{ // bishop1
				BishopFigure,
				3,
				3,
			},
			{ // king
				KingFigure,
				4,
				4,
			},
			{ // queen
				QueenFigure,
				5,
				5,
			},
			{ // bishop2
				BishopFigure,
				6,
				6,
			},
			{ // knight2
				KnightFigure,
				7,
				7,
			},
			{ // rook2
				RookFigure,
				8,
				8,
			},
		},
	}
	player2defaultFigurePositions = PlayerFigurePositions{
		[]FigureInfo{
			{ // pawn1
				PawnFigure,
				9,
				56,
			},
			{ // pawn2
				PawnFigure,
				10,
				55,
			},
			{ // pawn3
				PawnFigure,
				11,
				54,
			},
			{ // pawn4
				PawnFigure,
				12,
				53,
			},
			{ // pawn5
				PawnFigure,
				13,
				52,
			},
			{ // pawn6
				PawnFigure,
				14,
				51,
			},
			{ // pawn7
				PawnFigure,
				15,
				50,
			},
			{ // pawn8
				PawnFigure,
				16,
				49,
			},
			{ // rook1
				RookFigure,
				1,
				64,
			},
			{ // knight1
				KnightFigure,
				2,
				63,
			},
			{ // bishop1
				BishopFigure,
				3,
				62,
			},
			{ // king
				KingFigure,
				4,
				61,
			},
			{ // queen
				QueenFigure,
				5,
				60,
			},
			{ // bishop2
				BishopFigure,
				6,
				59,
			},
			{ // knight2
				KnightFigure,
				7,
				58,
			},
			{ // rook2
				RookFigure,
				8,
				57,
			},
		},
	}
}
