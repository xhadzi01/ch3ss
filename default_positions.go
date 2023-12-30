package main

var player1defaultFigurePositions PlayerFigurePositions
var player2defaultFigurePositions PlayerFigurePositions

func init() {
	player1defaultFigurePositions = PlayerFigurePositions{
		[]FigureInfo{
			{ // pawn1
				PawnFigure,
				9,
			},
			{ // pawn2
				PawnFigure,
				10,
			},
			{ // pawn3
				PawnFigure,
				11,
			},
			{ // pawn4
				PawnFigure,
				12,
			},
			{ // pawn5
				PawnFigure,
				13,
			},
			{ // pawn6
				PawnFigure,
				14,
			},
			{ // pawn7
				PawnFigure,
				15,
			},
			{ // pawn8
				PawnFigure,
				16,
			},
			{ // rook1
				RookFigure,
				1,
			},
			{ // knight1
				KnightFigure,
				2,
			},
			{ // bishop1
				BishopFigure,
				3,
			},
			{ // king
				KingFigure,
				4,
			},
			{ // queen
				QueenFigure,
				5,
			},
			{ // bishop2
				BishopFigure,
				6,
			},
			{ // knight2
				KnightFigure,
				7,
			},
			{ // rook2
				RookFigure,
				8,
			},
		},
	}
	player2defaultFigurePositions = PlayerFigurePositions{
		[]FigureInfo{
			{ // pawn1
				PawnFigure,
				56,
			},
			{ // pawn2
				PawnFigure,
				55,
			},
			{ // pawn3
				PawnFigure,
				54,
			},
			{ // pawn4
				PawnFigure,
				53,
			},
			{ // pawn5
				PawnFigure,
				52,
			},
			{ // pawn6
				PawnFigure,
				51,
			},
			{ // pawn7
				PawnFigure,
				50,
			},
			{ // pawn8
				PawnFigure,
				49,
			},
			{ // rook1
				RookFigure,
				64,
			},
			{ // knight1
				KnightFigure,
				63,
			},
			{ // bishop1
				BishopFigure,
				62,
			},
			{ // king
				KingFigure,
				61,
			},
			{ // queen
				QueenFigure,
				60,
			},
			{ // bishop2
				BishopFigure,
				59,
			},
			{ // knight2
				KnightFigure,
				58,
			},
			{ // rook2
				RookFigure,
				57,
			},
		},
	}
}
