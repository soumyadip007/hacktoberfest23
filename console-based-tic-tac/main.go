package main

import (
	"fmt"
)

var board [3][3]string
var currentPlayer string
var totalMoves int

func main() {
	initBoard()
	currentPlayer = "X"
	totalMoves = 0
	fmt.Println("Welcome to Tic-Tac-Toe!")

	for {
		printBoard()
		makeMove()
		if checkWin() {
			printBoard()
			fmt.Printf("Player %s wins!\n", currentPlayer)
			break
		} else if totalMoves == 9 {
			printBoard()
			fmt.Println("It's a draw!")
			break
		}
		switchPlayer()
	}
}

func initBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = " "
		}
	}
}

func printBoard() {
	fmt.Println("  0 1 2")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d ", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("%s", board[i][j])
			if j < 2 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if i < 2 {
			fmt.Println("  -+-+-")
		}
	}
	fmt.Println()
}

func makeMove() {
	var row, col int
	for {
		fmt.Printf("Player %s, enter row (0-2) and column (0-2) to place your %s (e.g., 1 2): ", currentPlayer, currentPlayer)
		_, err := fmt.Scanf("%d %d", &row, &col)
		if err != nil || row < 0 || row > 2 || col < 0 || col > 2 || board[row][col] != " " {
			fmt.Println("Invalid input. Try again.")
		} else {
			board[row][col] = currentPlayer
			totalMoves++
			break
		}
	}
}

func checkWin() bool {
	return checkRows() || checkColumns() || checkDiagonals()
}

func checkRows() bool {
	for i := 0; i < 3; i++ {
		if board[i][0] == currentPlayer && board[i][1] == currentPlayer && board[i][2] == currentPlayer {
			return true
		}
	}
	return false
}

func checkColumns() bool {
	for j := 0; j < 3; j++ {
		if board[0][j] == currentPlayer && board[1][j] == currentPlayer && board[2][j] == currentPlayer {
			return true
		}
	}
	return false
}

func checkDiagonals() bool {
	if (board[0][0] == currentPlayer && board[1][1] == currentPlayer && board[2][2] == currentPlayer) ||
		(board[0][2] == currentPlayer && board[1][1] == currentPlayer && board[2][0] == currentPlayer) {
		return true
	}
	return false
}

func switchPlayer() {
	if currentPlayer == "X" {
		currentPlayer = "O"
	} else {
		currentPlayer = "X"
	}
}

