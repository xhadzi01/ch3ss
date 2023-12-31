package main

// Chessboard layout
// |   ||  A |  B |  C |  D |  E |  F |  G |  H ||   |
// |--------------------------------------------------
// | 8 || 57 | 58 | 59 | 60 | 61 | 62 | 63 | 64 || 8 |
// | 7 || 49 | 50 | 51 | 52 | 53 | 54 | 55 | 56 || 7 |
// | 6 || 41 | 42 | 43 | 44 | 45 | 46 | 47 | 48 || 6 |
// | 5 || 33 | 34 | 35 | 36 | 37 | 38 | 39 | 40 || 5 |
// | 4 || 25 | 26 | 27 | 28 | 29 | 30 | 31 | 32 || 4 |
// | 3 || 17 | 18 | 19 | 20 | 21 | 22 | 23 | 24 || 3 |
// | 2 ||  9 | 10 | 11 | 12 | 13 | 14 | 15 | 16 || 2 |
// | 1 ||  1 |  2 |  3 |  4 |  5 |  6 |  7 |  8 || 1 |
// |--------------------------------------------------
// |   ||  A |  B |  C |  D |  E |  F |  G |  H ||   |

type FigureType uint64
type FigureIndex uint64

const (
	_ FigureType = iota
	PawnFigure
	RookFigure
	KnightFigure
	BishopFigure
	QueenFigure
	KingFigure
)

type FigureInfo struct {
	FigureType
	FigureIndex
	CurrentFigurePosition uint64
}

type PlayerFigurePositions struct {
	FigureInfoList []FigureInfo
}

type PlayerID string
type PlayerInfo struct {
	PlayerID
	PlayerFigurePositions
}

type PlayerType uint8

const (
	Player1 PlayerType = 0
	Player2 PlayerType = 1
)

func NewPlayerInfo(playerID PlayerID, pt PlayerType) PlayerInfo {
	if pt == Player1 {
		return PlayerInfo{
			PlayerID:              playerID,
			PlayerFigurePositions: player1defaultFigurePositions,
		}
	} else if pt == Player2 {
		return PlayerInfo{
			PlayerID:              playerID,
			PlayerFigurePositions: player2defaultFigurePositions,
		}
	}

	panic("unknown player type")
}

func NewPlayerFigurePositions(pt PlayerType) *PlayerFigurePositions {
	if pt == Player1 {
		retval := player1defaultFigurePositions
		return &retval
	} else if pt == Player2 {
		retval := player2defaultFigurePositions
		return &retval
	}

	panic("unknown player type")
}
