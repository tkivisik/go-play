// go-play is a package for playful exploration of Golang
package main

import (
	"flag"
	"fmt"

	"github.com/tkivisik/go-play/games"
)

func main() {
	var n = flag.Int("n", 1, "rounds to play")
	var game = flag.String("game", "", "game to play (e.g. 'num', 'bin')")
	flag.Parse()

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
