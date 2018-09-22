// go-play is a package for playful exploration of Golang
package main

import (
	"flag"
	"fmt"

	"github.com/tkivisik/go-play/games"
)

// PrintBoard prints the current battleship board
func PrintBoard() {
	fmt.Println("   a b c d e f g h")
	for i := 0; i < 8; i++ {
		fmt.Printf("%2d", i+1)
		for j := 0; j < 8; j++ {
			fmt.Print(" *")
		}
		fmt.Println()
	}
}

func main() {
	var n = flag.Int("n", 1, "rounds to play")
	var game = flag.String("game", "", "game to play (e.g. 'num', 'bin')")
	flag.Parse()

	PrintBoard()

	for i := 0; i < *n; i++ {
		switch *game {
		case "num":
			games.NumberGuessing()
		case "bin":
			games.BinaryGuessing()
		case "ship":
			games.Battleship()
		default:
			fmt.Println("Pass a flag -game with either 'num' or 'bin' or 'ship'")
		}
	}
}
