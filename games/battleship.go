package games

import (
	"crypto/rand"
	"fmt"
)

func PrintBoard() {
	for i := 0; i < 8; i++ {
		fmt.Println("* * * * * * * *")
	}
}

// Battleship is a game to guess the location of ships
func Battleship() {
}
