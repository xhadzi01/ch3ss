package main

import (
	"fmt"
	"math"

	"github.com/fatih/color"
)

var sepTemplate string = "  |-A---B---C---D---E---F---G---H-|\n"
var rowTemplate string = "%d | %s | %s | %s | %s | %s | %s | %s | %s |\n"

type colorFunction func(...interface{}) string

func translatePlayersPosition(figurePosition uint64) int {
	var figurePositionFlat int = 1
	for figurePosition != 1 {
		figurePositionFlat++
		figurePosition = figurePosition >> 1
	}
	figurePositionFlat--

	return figurePositionFlat
}

func fillInSingleFigure(values *[8][8]string, figurePosition uint64, playerFigureOutput PlayerFigureOutput, clrFunc colorFunction) {
	if figurePosition == 0 {
		// figure already down
		return
	}
	figurePositionFlat := translatePlayersPosition(figurePosition)

	var row int = figurePositionFlat / 8
	var col int = int(math.Mod(float64(figurePositionFlat), 8))
	aaa := fmt.Sprintf("%c", playerFigureOutput)
	values[row][col] = clrFunc(aaa)
}

func fillInFigures(values *[8][8]string, playerType PlayerType, playerFigures *PlayerFigures) {
	var colorForText func(...interface{}) string
	if playerType == Player1 {
		colorForText = color.New(color.FgBlue).SprintFunc()
	} else if playerType == Player2 {
		colorForText = color.New(color.FgRed).SprintFunc()
	} else {
		panic("Unknown player type")
	}

	fillInSingleFigure(values, playerFigures.pawn1, pawnFigureOutput, colorForText)
	fillInSingleFigure(values, playerFigures.pawn2, pawnFigureOutput, colorForText)
	fillInSingleFigure(values, playerFigures.pawn3, pawnFigureOutput, colorForText)
	fillInSingleFigure(values, playerFigures.pawn4, pawnFigureOutput, colorForText)
	fillInSingleFigure(values, playerFigures.pawn5, pawnFigureOutput, colorForText)
	fillInSingleFigure(values, playerFigures.pawn6, pawnFigureOutput, colorForText)
	fillInSingleFigure(values, playerFigures.pawn7, pawnFigureOutput, colorForText)
	fillInSingleFigure(values, playerFigures.pawn8, pawnFigureOutput, colorForText)
	fillInSingleFigure(values, playerFigures.rook1, rookFigureOutput, colorForText)
	fillInSingleFigure(values, playerFigures.knight1, knightFigureOutput, colorForText)
	fillInSingleFigure(values, playerFigures.bishop1, bishopFigureOutput, colorForText)
	fillInSingleFigure(values, playerFigures.queen, queenFigureOutput, colorForText)
	fillInSingleFigure(values, playerFigures.king, kingFigureOutput, colorForText)
	fillInSingleFigure(values, playerFigures.bishop2, bishopFigureOutput, colorForText)
	fillInSingleFigure(values, playerFigures.knight2, knightFigureOutput, colorForText)
	fillInSingleFigure(values, playerFigures.rook2, rookFigureOutput, colorForText)
}

func translateFigurePositions(board *Board) [8][8]string {
	var values [8][8]string

	// clear the board
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			values[i][j] = " "
		}
	}

	// fill in player1 positions
	fillInFigures(&values, Player1, board.player1)

	// fill in player2 positions
	fillInFigures(&values, Player2, board.player2)

	return values
}

func printBoardToScreen(board *Board) {
	var valuesInRows [8][8]string = translateFigurePositions(board)
	var outString string = sepTemplate
	for i := 7; i >= 0; i-- {
		outString += fmt.Sprintf(rowTemplate, i, valuesInRows[i][0], valuesInRows[i][1], valuesInRows[i][2], valuesInRows[i][3], valuesInRows[i][4], valuesInRows[i][5], valuesInRows[i][6], valuesInRows[i][7])
	}
	outString += sepTemplate
	fmt.Println(outString)
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
