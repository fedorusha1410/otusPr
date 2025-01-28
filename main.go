package main

import (
	"fmt"
)

func GetChessBoard(size int) string {
	if size <= 0 {
		fmt.Println("Size cannot be less or equal 0")
	}

	white := " "
	black := "#"
	result := ""
	emptyMatrix := make([][]string, size)

	for i := range emptyMatrix {
		emptyMatrix[i] = make([]string, size)
	}

	for i := range size {
		for j := range size {
			if (i+j)%2 == 0 {
				emptyMatrix[i][j] = white
			} else {
				emptyMatrix[i][j] = black
			}

			result += emptyMatrix[i][j]
		}

		result += "\n"
	}

	return result

}

func main() {

	board := GetChessBoard(8)
	fmt.Print(board)

}
