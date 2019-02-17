// Copyright © 2018 Taavi Kivisik
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package games

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

const (
	letters         string = "abcdefghijklmnopqrstuvwxyz"
	boardSideLength uint8  = 4
	MaxShips        int    = 2
	maxShots        int    = 100
)

// Legend holds the mapping of a board
type LegendStruct struct {
	Terrain string
	Ship    string
	Hit     string
	Miss    string
}

var Legend = LegendStruct{
	Terrain: "~",
	Ship:    "0",
	Hit:     "X",
	Miss:    "*",
}

// String returns a printable string of a Legend
func (l *LegendStruct) String() string {
	var str strings.Builder

	fmt.Fprintf(&str, "%10s '%s'\n", "Terrain", l.Terrain)
	fmt.Fprintf(&str, "%10s '%s'\n", "Ship", l.Ship)
	fmt.Fprintf(&str, "%10s '%s'\n", "Hit", l.Hit)
	fmt.Fprintf(&str, "%10s '%s'\n", "Miss", l.Miss)

	return str.String()
}

// Coordinate expresses a location on a map using x and y
type Coordinate struct {
	x byte
	y byte
}

func (c *Coordinate) String() string {
	return fmt.Sprintf("Human representation: %c%d", byte(letters[c.x]), c.y+1)
}

// Read prompts the user to enter a Coordinate
func (c *Coordinate) Read() {
	var s string
	for i := 0; i < 100; i++ {
		fmt.Println("Please enter a coordinate (e.g. 'd3'):")
		fmt.Scanln(&s)
		if s == "" {
			continue
		}
		if strings.Contains(letters[boardSideLength:], string(s[0])) {
			fmt.Printf("Please use letters from %c-%c\n", letters[0], letters[int(boardSideLength)])
			continue
		}
		if strings.Contains("0123456789", string(s[0])) {
			fmt.Printf("First character must be a letter from %c-%c\n", letters[0], letters[int(boardSideLength)])
			continue
		}
		y, err := strconv.Atoi(s[1:])
		if err != nil {
			fmt.Println("Please make sure number follows the letter immediately")
			continue
		}
		if y <= 0 || uint8(y) > boardSideLength {
			fmt.Printf("Please use numbers from 1-%d\n", boardSideLength)
			continue
		}
		c.x = byte(strings.IndexRune(letters, rune(s[0])))
		c.y = byte(y)
		c.y--
		break
	}
}

type layer [boardSideLength]byte

func (l *layer) StringRaw() string {
	var str strings.Builder

	for row := uint8(0); row < boardSideLength; row++ {
		fmt.Fprintf(&str, "%08b\n", l[row])
	}
	return str.String()
}

// coordinateToOne turns a Coordinate on a layer to 1
func (l *layer) coordinateToOne(c Coordinate) {
	l[c.y] |= 1 << c.x
}

// Board is an object for ships and shots
type Board struct {
	ships    layer
	shots    layer
	HitCount int8
}

func NewBoard() *Board {
	return &Board{ships: layer{}, shots: layer{}, HitCount: 0}
}

// String returns the current battleship board as string
func (b *Board) String(enemy bool) string {
	b.HitCount = 0
	var str strings.Builder

	str.WriteString("  ") // space instead of a row number
	for column := uint8(0); column < boardSideLength; column++ {
		fmt.Fprintf(&str, "%2s", letters[column:column+1])
	}
	str.WriteString("\n")

	for row := uint8(0); row < boardSideLength; row++ {

		fmt.Fprintf(&str, "%2d", row+1)
		for column := uint8(0); column < boardSideLength; column++ {
			coord := Coordinate{column, row}
			if b.hasShip(coord) {
				if b.hasShot(coord) {
					fmt.Fprintf(&str, " %s", Legend.Hit)
					b.HitCount++
				} else {
					if enemy { // hide enemy ships until hit
						fmt.Fprintf(&str, " %s", Legend.Terrain)
					} else {
						fmt.Fprintf(&str, " %s", Legend.Ship)
					}
				}
			} else { // not a ship
				if b.hasShot(coord) {
					fmt.Fprintf(&str, " %s", Legend.Miss)
				} else {
					fmt.Fprintf(&str, " %s", Legend.Terrain)
				}
			}
		}
		str.WriteString("\n")
	}
	return str.String()
}

// Print prints the board according to the Legend
func (b *Board) Print(enemy bool) {
	fmt.Print(b.String(enemy))
}

func (b *Board) hasShip(c Coordinate) bool {
	return b.ships[c.y]&(1<<c.x) != 0
}

func (b *Board) hasShot(c Coordinate) bool {
	return b.shots[c.y]&(1<<c.x) != 0
}

func (b *Board) isSurroundedByWater(c Coordinate) bool {
	for row := -1; row < 2; row++ {
		if int(c.y)+row < 0 {
			continue
		}
		if c.y+byte(row) >= boardSideLength {
			return true
		}
		for col := -1; col < 2; col++ {
			if int(c.x)+col < 0 {
				continue
			}
			if c.x+byte(col) >= boardSideLength {
				continue
			}

			if b.ships[c.y+byte(row)]&(1<<(c.x+byte(col))) != 0 {
				return false
			}
		}
	}
	return true
}

func (b *Board) randomLocation() Coordinate {
	var coord Coordinate
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(boardSideLength*boardSideLength)))
	if err != nil {
		panic(err)
	}
	n := uint8(nBig.Int64())
	if err != nil {
		fmt.Println("error:", err)
	}
	coord.x = n % boardSideLength
	coord.y = n / boardSideLength

	return coord
}

func (b *Board) AddShipBy(random bool) {
	var coord Coordinate

	if random {
		for {
			coord = b.randomLocation()
			if b.hasShip(coord) == true {
				continue
			}
			if b.isSurroundedByWater(coord) == false {
				continue
			}
			b.ships.coordinateToOne(coord)
			break
		}
	} else {
		for {
			coord.Read()
			if b.hasShip(coord) == true {
				fmt.Println("There already is a ship in that location.")
				continue
			}
			if b.isSurroundedByWater(coord) == false {
				fmt.Println("Ships must have space between them.")
				continue
			}
			b.ships.coordinateToOne(coord)
			break
		}
	}
}

func (b *Board) ShootThisBoard(automatic bool) {
	var coord Coordinate

	if automatic == true {
		for i := 0; i < 1000; {
			coord = b.randomLocation()
			if b.hasShot(coord) {
				continue
			}
			b.shots.coordinateToOne(coord)
			fmt.Println("ENEMY just shot.")
			break
		}
	} else {
		for i := 0; i < 1000; {
			coord.Read()
			if b.hasShot(coord) {
				continue
			}
			b.shots.coordinateToOne(coord)
			break
		}
	}
}
