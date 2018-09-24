// go-play is a package for playful exploration of Golang
package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/tkivisik/go-play/games"
)

// BoardSideLength determines the height and width of the game board
const BoardSideLength uint8 = 8

// Letters hold the alphabet
const Letters string = "abcdefghijklmnopqrstuvwxyz"

// MaxShips determiens the number of ships per player
const MaxShips int = 2

// MaxShots determines the maximum length of a game
const MaxShots int = 100

// Legend is a dictionary of program logic to map elements
var Legend = map[string]string{
	"terrain": "~",
	"ship":    "0",
	"hit":     "X",
	"miss":    "*",
}

// PrintLegend prints the legend on screen
func PrintLegend() {
	fmt.Println("")
	for k, v := range Legend {
		fmt.Printf("%2s - %s    ", v, k)
	}
	fmt.Println("\n Valid point example: 'd 3'")
	fmt.Print("\n")
}

// PrintBoard prints the current battleship board
func PrintBoard(ships, shots map[int8]byte) int {
	hitCount := 0
	fmt.Print("\n  ")
	for i := uint8(0); i < BoardSideLength; i++ {
		fmt.Printf("%2s", Letters[i:i+1])
	}
	fmt.Println()

	for i := uint8(0); i < BoardSideLength; i++ {
		fmt.Printf("%2d", i+1)
		for j := uint8(0); j < BoardSideLength; j++ {
			if ships[int8(i)]&(1<<j) != 0 {
				if shots[int8(i)]&(1<<j) != 0 {
					fmt.Printf(" %s", Legend["hit"])
					hitCount++
				} else {
					fmt.Printf(" %s", Legend["ship"])
				}
			} else {
				if shots[int8(i)]&(1<<j) != 0 {
					fmt.Printf(" %s", Legend["miss"])
				} else {
					fmt.Printf(" %s", Legend["terrain"])
				}
			}
		}
		fmt.Println()
	}
	fmt.Println()
	return hitCount
}

// Shoot updates the shots map
func PointToMap(point Point, landscape map[int8]byte) {
	point.number--
	var letterIndex uint8 = uint8(strings.Index(Letters, point.letter))
	landscape[int8(point.number)] |= 1 << uint8(letterIndex)
}

type Point struct {
	letter string
	number int
}

// PlaceShips guides the player to place its ships
func PlaceShips() map[int8]byte {
	var point Point
	ships := map[int8]byte{0: uint8(0)}
	shots := map[int8]byte{}

	fmt.Printf("Please choose the location of your %d ships (e.g. 'd 3')\n", MaxShips)
	for i := 0; i < MaxShips; i++ {
		PrintBoard(ships, shots)
		fmt.Scanln(&point.letter, &point.number)
		PointToMap(point, ships)
	}
	PrintBoard(ships, shots)
	return ships
}

func PlaceShots(ships map[int8]byte) {
	var point Point
	var Scoreboard = map[string]int{
		"me":    0,
		"enemy": 0,
	}
	shots := map[int8]byte{}

	fmt.Print("Please choose where to shoot (e.g. 'd 2')\n\n")
	for i := 0; i < MaxShots; i++ {
		fmt.Scanln(&point.letter, &point.number)
		PointToMap(point, shots)
		PrintLegend()
		Scoreboard["me"] = PrintBoard(ships, shots)
		if Scoreboard["me"] >= MaxShips {
			fmt.Println(" ALL SHIPS SUNK! WELL DONE!")
			break
		}
	}
}

func main() {
	var n = flag.Int("n", 1, "rounds to play")
	var game = flag.String("game", "", "game to play (e.g. 'num', 'bin')")
	flag.Parse()

	fmt.Println(" BATTLESHIP THE GAME\n")

	ships := PlaceShips()
	PlaceShots(ships)

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
