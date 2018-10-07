package games

import (
	"crypto/rand"
	"fmt"
	"os"
	"strings"
)

// boardSideLength determines the height and width of the game board
// Letters hold the alphabet
// maxShips determiens the number of ships per player
// MaxShots determines the maximum length of a game
const (
	boardSideLength uint8  = 4
	letters         string = "abcdefghijklmnopqrstuvwxyz"
	maxShips        int    = 2
	maxShots        int    = 100
)

// legend is a dictionary of program logic to map elements
var legend = map[string]string{
	"terrain": "~",
	"ship":    "0",
	"hit":     "X",
	"miss":    "*",
}

// printLegend prints the legend on screen
func printLegend() {
	fmt.Println("")
	for k, v := range legend {
		fmt.Printf("%2s - %s    ", v, k)
	}
	fmt.Println("\n Valid coordinate example: 'd 3'")
	fmt.Print("\n")
}

// printBoard prints the current battleship board
func printBoard(ships, shots map[int8]byte, enemy bool) int {
	hitCount := 0

	fmt.Print("\n  ")
	for column := uint8(0); column < boardSideLength; column++ {
		fmt.Printf("%2s", letters[column:column+1])
	}
	fmt.Println()

	for row := uint8(0); row < boardSideLength; row++ {
		fmt.Printf("%2d", row+1)
		for column := uint8(0); column < boardSideLength; column++ {
			if ships[int8(row)]&(1<<column) != 0 { // ship found
				if shots[int8(row)]&(1<<column) != 0 { // shot found
					fmt.Printf(" %s", legend["hit"])
					hitCount++
				} else {
					if enemy { // hide enemy ships until hit
						fmt.Printf(" %s", legend["terrain"])
					} else {
						fmt.Printf(" %s", legend["ship"])
					}
				}
			} else { // not a ship
				if shots[int8(row)]&(1<<column) != 0 { // shot found
					fmt.Printf(" %s", legend["miss"])
				} else {
					fmt.Printf(" %s", legend["terrain"])
				}
			}
		}
		fmt.Println()
	}
	fmt.Println()
	return hitCount
}

// coordinateToMap updates the shots map
func coordinateToMap(p coordinate, landscape map[int8]byte) {
	c.y--
	var letterIndex = uint8(strings.Index(letters, c.x))
	landscape[int8(c.y)] |= 1 << uint8(letterIndex)
}

// coordinate expresses a location on a map as a
//type coordinate struct {
//	x int
//	y int
//}

// PlaceShips guides the player to place its ships
func placeShips(random bool) map[int8]byte {
	var p coordinate
	ships := map[int8]byte{0: uint8(0)}
	shots := map[int8]byte{}

	if random {
		for nShipsPlaced := 0; nShipsPlaced < maxShips; nShipsPlaced++ {
			b := make([]byte, 2)
			_, err := rand.Read(b)

			if err != nil {
				fmt.Println("error:", err)
			}
			c.x = letters[int(b[0]%boardSideLength) : int(b[0]%boardSideLength)+1]
			c.y = int(b[1] % boardSideLength)
			c.y++ // Convert into a number from 1 to boardSideLength

			//fmt.Printf("%d, %d", b[0]%boardSideLength, b[1]%boardSideLength)
			coordinateToMap(p, ships)
		}

	} else {
		fmt.Printf("Please choose the location of your %d ships (e.g. 'd 3')\n", maxShips)
		for nShipsPlaced := 0; nShipsPlaced < maxShips; nShipsPlaced++ {
			fmt.Scanln(&c.x, &c.y)
			coordinateToMap(p, ships)
			printBoard(ships, shots, false)
		}
	}
	return ships
}

func placeShot(ships, shots map[int8]byte, automatic bool) {
	var p coordinate

	if automatic {
		b := make([]byte, 2)
		_, err := rand.Read(b)
		if err != nil {
			fmt.Println("error:", err)
		}

		c.x = letters[int(b[0]%boardSideLength) : int(b[0]%boardSideLength)+1]
		c.y = int(b[1] % boardSideLength)
		c.y++ // Convert into a number from 1 to boardSideLength

		coordinateToMap(p, shots)
	} else {
		fmt.Print("Please choose where to shoot (e.g. 'd 2')\n\n")
		fmt.Scanln(&c.x, &c.y)
		coordinateToMap(p, shots)
		printLegend()
	}
}

// Battleship is a game to guess the location of ships
func Battleship() {
	fmt.Println("***********************")
	fmt.Println("* BATTLESHIP THE GAME *")
	fmt.Print("***********************\n\n")

	var enemyScore, myScore int
	enemyShips := placeShips(true)
	enemyShots := map[int8]byte{0: uint8(0)}
	myShips := placeShips(false)
	myShots := map[int8]byte{0: uint8(0)}

	for round := int8(0); int(round) < maxShots; round++ {
		fmt.Println("ENEMY MOVE:")
		fmt.Println("MY BOARD, ENEMY SHOTS:")
		placeShot(myShips, enemyShots, true)
		enemyScore = printBoard(myShips, enemyShots, false)

		fmt.Println("\t\t\tMY MOVE:")
		fmt.Println("\t\t\tENEMY BOARD, MY SHOTS:")
		placeShot(enemyShips, myShots, false)
		myScore = printBoard(enemyShips, myShots, true)

		if enemyScore >= maxShips {
			if myScore >= maxShips {
				fmt.Println("IT'S A DRAW, GAME OVER, WELL DONE")
			} else {
				fmt.Println("GAME OVER, YOU LOST")
			}
			os.Exit(0)
		} else {
			if myScore >= maxShips {
				fmt.Println("GAME OVER. YOU WON !!!")
				os.Exit(0)
			}
		}

		fmt.Println("===== ===== ===== ===== ===== ===== ===== ===== ===== =====")
		fmt.Println("   ===== ===== ===== ===== ===== ===== ===== ===== =====   ")
	}
}
