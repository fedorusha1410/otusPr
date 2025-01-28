package main

import (
	"fmt"
	"strings"
)

func GetChessBoard(size int) string {
	if size <= 0 {
		return "Size cannot be less or equal 0"
	}

	white := " "
	black := "#"
	var result strings.Builder

	for i := range size {
		for j := range size {
			if (i+j)%2 == 0 {
				result.WriteString(white)
			} else {
				result.WriteString(black)
			}
		}

		result.WriteString("\n")
	}

	return result.String()

}

func main() {

	var size int
	fmt.Println("Enter the size of chess board: ")
	fmt.Scanf("%d", &size)
	fmt.Printf("You wrote %d\n", size)
	board := GetChessBoard(size)
	fmt.Println(board)

}
