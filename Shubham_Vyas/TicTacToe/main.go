package main

import (
	"fmt"
	"strconv"
	"strings"
)

// takes no of columns in tic tac toe
func createPatterns(cols int) (patterns [][]int) {
	// to win you must satisfy one of the following
	// First/Second/Third row is all 1 or 0
	// First/Second/Third col is all 1 or 0
	// Diagonals must be all 1 or 0
	totalCells := cols * cols

	// rows
	for i := 0; i < totalCells; i += cols {
		rows := []int{}
		for j := i; j < i+cols; j++ {
			rows = append(rows, j)
		}
		patterns = append(patterns, rows)
	}
	// cols
	for i := 0; i < cols; i++ {
		columns := []int{}
		for j := i; j < totalCells; j += cols {
			columns = append(columns, j)
		}
		patterns = append(patterns, columns)
	}
	// left diagonals
	leftDiagonal := []int{}
	for j := 0; j < totalCells; j += cols + 1 {
		leftDiagonal = append(leftDiagonal, j)
	}
	patterns = append(patterns, leftDiagonal)
	// right dialog
	rightDiagonal := []int{}
	for i := cols - 1; len(rightDiagonal) < cols; i += cols - 1 {
		rightDiagonal = append(rightDiagonal, i)
	}
	patterns = append(patterns, rightDiagonal)
	return
}

func returnGridToPrint(inputs []int, cols int) (x string) {
	for i, v := range inputs {
		if i%cols == 0 && i > 0 {
			x += "\n" + strings.Repeat("---", cols) + "\n"
		}
		if v == 0 {
			x += " "
		} else if v > 0 {
			x += "X"
		} else {
			x += "O"
		}
		if (i+1)%cols != 0 {
			x += " | "
		}
	}
	return
}

func main() {
	// Select grid
	var columns int
	for {
		columnInput := ""
		fmt.Printf("Please enter grid between 3 to 10: ")
		fmt.Scanln(&columnInput)

		var err error

		columns, err = strconv.Atoi(columnInput)
		if err != nil || columns < 3 || columns > 10 {
			fmt.Println("Invalid grid selected")
			continue
		}
		break
	}

	patterns := createPatterns(columns)

	// 1 represent X and 0 represent O
	// key would be position and value will be 1/0
	entries := make(map[int]int)

	// total cells in grid
	// 3x3 grid will have 9 cells
	totalCells := columns * columns

	playerInputs := make([]int, totalCells, totalCells)
	player := 1
	// to know which players chance it is
	for {
		// Print tic-tac-toe
		fmt.Println()
		fmt.Println(returnGridToPrint(playerInputs, columns))
		fmt.Printf("Enter a value between 1-9: ")
		var scannedInput string
		fmt.Scanln(&scannedInput)

		input, err := strconv.Atoi(scannedInput)
		// if conversion fail instead of exiting ask for a new value
		if err != nil || input == 0 || input > 9 {
			continue
		}
		// if input is already selected
		if _, ok := entries[input]; ok {
			continue
		}
		entries[input] = 1
		// for even its O and or odd it will be X
		if player%2 == 1 {
			playerInputs[input-1] = 1
		} else {
			playerInputs[input-1] = -1
		}
		player += 1
		// To win you must have one of patterns
		anyoneWon := false

		for _, pattern := range patterns {
			sum := 0
			for _, v := range pattern {
				sum += playerInputs[v]
			}
			if sum == columns || 0 == columns+sum {
				if sum > 0 {
					fmt.Println("Player 1 wins")
				} else {
					fmt.Println("Player 2 wins")
				}
				anyoneWon = true
				break
			}
		}
		if anyoneWon {
			break
		}
	}
}
