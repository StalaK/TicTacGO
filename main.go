package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	clearScreen()
	fmt.Println("  ____    _         _____            _____  _____")
	fmt.Println(" | _ _|  |_|   ___ |_   _| ___  ___ |   __||     |")
	fmt.Println("  | |    | |  |  _|  | |  | .'||  _||  |  ||  |  |")
	fmt.Println("  |_|    |_|  |___|  |_|  |__,||___||_____||_____|")

	fmt.Println("\nPress enter to begin")
	fmt.Scanln()
	clearScreen()

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter Player 1 name: ")
	scanner.Scan()
	xPlayer := scanner.Text()

	fmt.Print("Enter Player 2 name: ")
	scanner.Scan()
	oPlayer := scanner.Text()

	var playGame bool = true
	var xPlayerScore uint8 = 0
	var oPlayerScore uint8 = 0
	var gameCount uint8 = 0

	for playGame == true {
		gameCount++
		xPlayerScore, oPlayerScore = newGame(xPlayer, oPlayer, xPlayerScore, oPlayerScore, gameCount)

		fmt.Println("\nPlay another game? (true/false)")
		fmt.Scan(&playGame)
	}

	var winnerText string
	if xPlayerScore > oPlayerScore {
		winnerText = fmt.Sprint(xPlayer, " wins!")
	} else if xPlayerScore < oPlayerScore {
		winnerText = fmt.Sprint(oPlayer, " wins!")
	} else {
		winnerText = "The game ended in a draw!"
	}

	fmt.Println("\nThank you for playing!")
	fmt.Println("\nThe final score was\n\t", xPlayer, "\t:\t", xPlayerScore, "\n\t", oPlayer, "\t:\t", oPlayerScore)
	fmt.Println("\n", winnerText)
	fmt.Println("\n\nPress the enter key to exit")
	fmt.Scan()
}

func newGame(xPlayer, oPlayer string, xPlayerScore, oPlayerScore, gameCount uint8) (uint8, uint8) {
	board := [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	printGrid(board, xPlayer, oPlayer, xPlayerScore, oPlayerScore, gameCount)

	gameOver := false
	var turnCount uint8 = 0
	var winner int8 = -1

	for gameOver == false && turnCount < 9 {
		playerTurn := (gameCount + turnCount) % 2
		validMove := false

		for validMove == false {
			board, validMove, gameOver, winner = takeMove(playerTurn, xPlayer, oPlayer, board)
			if validMove == false {
				fmt.Println("Invalid position. Please try again")
			}
		}

		printGrid(board, xPlayer, oPlayer, xPlayerScore, oPlayerScore, gameCount)
		turnCount++
	}

	if winner == 0 {
		fmt.Println(xPlayer, " wins!")
		xPlayerScore++
	} else if winner == 1 {
		fmt.Println(oPlayer, " wins!")
		oPlayerScore++
	} else {
		fmt.Println("Draw!")
	}
	return xPlayerScore, oPlayerScore
}

func takeMove(playerTurn uint8, xPlayer, oPlayer string, board [9]string) ([9]string, bool, bool, int8) {
	var symbol string
	if playerTurn == 0 {
		fmt.Print(oPlayer)
		symbol = "O"
	} else {
		fmt.Print(xPlayer)
		symbol = "X"
	}

	var enteredPos string = "1"
	fmt.Print(" pick your position: ")
	fmt.Scan(&enteredPos)

	validMove := false
	gameOver := false
	var winningSymbol string
	var winner int8 = -1

	validGuess, pos := validGuess(enteredPos)
	if !validGuess {
		return board, validMove, gameOver, winner
	}

	if board[pos-1] != "X" && board[pos-1] != "O" {
		validMove = true
		board[pos-1] = symbol

		gameOver, winningSymbol = checkBoard(board)
	}

	if winningSymbol == "X" {
		winner = 0
	} else if winningSymbol == "O" {
		winner = 1
	}

	return board, validMove, gameOver, winner
}

func checkBoard(board [9]string) (bool, string) {
	// Horizontals
	if board[0] == board[1] && board[0] == board[2] {
		return true, board[0]
	}

	if board[3] == board[4] && board[3] == board[5] {
		return true, board[0]
	}

	if board[6] == board[7] && board[6] == board[8] {
		return true, board[0]
	}

	// Verticals
	if board[0] == board[3] && board[0] == board[6] {
		return true, board[0]
	}

	if board[1] == board[4] && board[1] == board[7] {
		return true, board[0]
	}

	if board[2] == board[5] && board[2] == board[8] {
		return true, board[0]
	}

	// Diagonals
	if board[0] == board[4] && board[0] == board[8] {
		return true, board[0]
	}

	if board[2] == board[4] && board[2] == board[6] {
		return true, board[0]
	}

	return false, ""
}

func printGrid(board [9]string, xPlayer, oPlayer string, xPlayerScore, oPlayerScore, gameCount uint8) {
	clearScreen()
	fmt.Println("X:", xPlayer, "(", xPlayerScore, ") vs. O:", oPlayer, "(", oPlayerScore, ") -- Game ", gameCount)

	fmt.Println("\n        |       |")
	fmt.Println("  ", board[0], "   |  ", board[1], "  |  ", board[2])
	fmt.Println("________|_______|________")
	fmt.Println("        |       |")
	fmt.Println("  ", board[3], "   |  ", board[4], "  |  ", board[5])
	fmt.Println("________|_______|________")
	fmt.Println("        |       |")
	fmt.Println("  ", board[6], "   |  ", board[7], "  |  ", board[8])
	fmt.Println("        |       |")
}

func clearScreen() {
	// Windows Only - Add Linux/Unix support
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func validGuess(guess string) (bool, int64) {
	parsedGuess, err := strconv.ParseInt(guess, 10, 8)
	if err != nil {
		return false, 0
	}

	if parsedGuess > 0 && parsedGuess < 10 {
		return true, parsedGuess
	}

	return false, 0
}
