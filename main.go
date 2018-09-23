// go-play is a package for playful exploration of Golang
package main

import (
	"flag"
	"fmt"

	"github.com/tkivisik/go-play/games"
)

func PrintLegend() {
	fmt.Println(" . - unhit location    0 - ship    x - miss    X - a hit")
	fmt.Println("")
}

// PrintBoard prints the current battleship board
func PrintBoard(ships, shots map[int8]byte) {
	fmt.Println("   a b c d e f g")
	for i := int8(1); i <= 7; i++ {
		fmt.Printf("%2d", i)
		for j := uint8(1); j <= 7; j++ {
			if ships[i]&(1<<j) != 0 {
				if shots[i]&(1<<j) != 0 {
					fmt.Print(" X")
				} else {
					fmt.Print(" 0")
				}
			} else {
				if shots[i]&(1<<j) != 0 {
					fmt.Print(" x")
				} else {
					fmt.Print(" .")
				}
			}
		}
		fmt.Println()
	}
}

// Shoot updates the shots map
func Shoot(point Point, shots map[int8]byte) {
	shots[int8(point.x)] |= 1 << uint8(point.y)
}

type Point struct{ x, y int }

func main() {
	var n = flag.Int("n", 1, "rounds to play")
	var game = flag.String("game", "", "game to play (e.g. 'num', 'bin')")
	flag.Parse()

	var point Point

	ships := map[int8]byte{
		1: 1 << uint8(1),
		2: 1 << uint8(4)}
	shots := map[int8]byte{}
	/*	shots := map[int8]byte{
		1: 1 << uint8(7),
		3: 1 << uint8(4)} */

	PrintLegend()
	PrintBoard(ships, shots)
	fmt.Println()

	for i := 0; i < 5; i++ {
		fmt.Scanln(&point.x, &point.y)
		Shoot(point, shots)
		PrintBoard(ships, shots)
	}

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
