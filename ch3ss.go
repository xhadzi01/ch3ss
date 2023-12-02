package main

import (
	"fmt"
	"time"
)

func main() {
	var board *Board = NewBoard()

	var i int = 0
	for {
		clearScreen()
		printBoardToScreen(board)
		fmt.Println(i)
		i++

		time.Sleep(time.Second * 1)
	}

}
