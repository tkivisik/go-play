package games

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

const (
	letters         string = "abcdefghijklmnopqrstuvwxyz"
	boardSideLength uint8  = 4
	maxShips        int    = 2
	maxShots        int    = 100
)

// Legend holds the mapping of a board
type Legend struct {
	Terrain string
	Ship    string
	Hit     string
	Miss    string
}

var legend = Legend{
	Terrain: "~",
	Ship:    "0",
	Hit:     "X",
	Miss:    "*",
}

func (l *Legend) String() string {
	var str strings.Builder

	fmt.Fprintf(&str, "%10s '%s'\n", "Terrain", l.Terrain)
	fmt.Fprintf(&str, "%10s '%s'\n", "Ship", l.Ship)
	fmt.Fprintf(&str, "%10s '%s'\n", "Hit", l.Hit)
	fmt.Fprintf(&str, "%10s '%s'\n", "Miss", l.Miss)

	return str.String()
}

// Coordinate expresses a location on a map as a
type Coordinate struct {
	x byte
	y byte
}

func (c *Coordinate) String() string {
	return fmt.Sprintf("Human representation: %c%d", byte(letters[c.x]), c.y+1)
}

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

type layer struct {
	layer [boardSideLength]byte
}

func (l *layer) StringRaw() string {
	var str strings.Builder

	for row := uint8(0); row < boardSideLength; row++ {
		fmt.Fprintf(&str, "%08b\n", l.layer[row])
	}
	return str.String()
}

func (l *layer) coordinateToOne(c Coordinate) {
	l.layer[c.y] |= 1 << c.x
}

// Board is an object for ships and shots
type Board struct {
	ships    layer
	shots    layer
	hitCount int8
}

// String returns the current battleship board as string
func (b *Board) String(enemy bool) string {
	b.hitCount = 0
	var str strings.Builder

	str.WriteString("\n  ")
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
					fmt.Fprintf(&str, " %s", legend.Hit)
					b.hitCount++
				} else {
					if enemy { // hide enemy ships until hit
						fmt.Fprintf(&str, " %s", legend.Terrain)
					} else {
						fmt.Fprintf(&str, " %s", legend.Ship)
					}
				}
			} else { // not a ship
				if b.hasShot(coord) {
					fmt.Fprintf(&str, " %s", legend.Miss)
				} else {
					fmt.Fprintf(&str, " %s", legend.Terrain)
				}
			}
		}
		str.WriteString("\n")
	}
	str.WriteString("\n")
	return str.String()
}

// Print prints the board according to the legend
func (b *Board) Print(enemy bool) {
	fmt.Print(b.String(enemy))
}

func (b *Board) hasShip(c Coordinate) bool {
	return b.ships.layer[c.y]&(1<<c.x) != 0
}

func (b *Board) hasShot(c Coordinate) bool {
	return b.shots.layer[c.y]&(1<<c.x) != 0
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

func (b *Board) init(random bool) {
	var coord Coordinate

	if random {
		for nShipsPlaced := 0; nShipsPlaced < maxShips; {
			coord = b.randomLocation()
			if b.hasShip(coord) == true {
				continue
			}
			b.ships.coordinateToOne(coord)
			nShipsPlaced++
		}
	} else {
		b.Print(false)
		fmt.Printf("Please select the location for your %d ships.\n", maxShips)
		for nShipsPlaced := 0; nShipsPlaced < maxShips; {
			coord.Read()
			if b.hasShip(coord) == true {
				fmt.Println("There already is a ship in that location.")
				continue
			}
			b.ships.coordinateToOne(coord)
			nShipsPlaced++
			b.Print(false)
		}
	}
}

func (b *Board) shootThisBoard(automatic bool) {
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

// Battleship is a game to guess the location of ships
func Battleship() {
	var str strings.Builder
	title := "BATTLESHIP THE GAME"

	str.WriteString(strings.Repeat("*", len(title)+4))
	fmt.Fprintf(&str, "\n* %s *\n", title)
	fmt.Fprintf(&str, "%s\n", strings.Repeat("*", len(title)+4))
	str.WriteString(legend.String())
	fmt.Print(str.String())
	str.Reset()

	var enemy, me Board
	var myScore, enemyScore int

	enemy.init(true)
	me.init(false)
	me.Print(false)

	for round := int8(0); int(round) < maxShots; round++ {
		str.WriteString(strings.Repeat("-=|=- ", 10))
		fmt.Fprintf(&str, "\n Round %2d\n", int(round)+1)
		fmt.Print(str.String())
		str.Reset()

		me.shootThisBoard(true)
		enemy.shootThisBoard(false)

		str.WriteString(strings.Repeat("\n", 20))
		fmt.Fprintf(&str, " %s%s\t\t %s", "ME", strings.Repeat(" ", int(boardSideLength)*2), "ENEMY")
		fmt.Print(str.String())
		str.Reset()

		meStrParts := strings.Split(me.String(false), "\n")
		enemyStrParts := strings.Split(enemy.String(true), "\n")
		for i := 0; i < len(meStrParts); i++ {
			fmt.Printf("%s\t\t%s\n", meStrParts[i], enemyStrParts[i])
		}

		myScore = int(enemy.hitCount)
		enemyScore = int(me.hitCount)

		fmt.Printf("SCORE :: Me: %d\t\tEnemy:%d\n", myScore, enemyScore)

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
	}
}
