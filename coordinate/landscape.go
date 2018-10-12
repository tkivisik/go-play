package coordinate

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	letters         string = "abcdefghijklmnopqrstuvwxyz"
	boardSideLength uint8  = 4
)

// Legend holds the mapping of a board
type Legend struct {
	Terrain rune
	Ship    byte
	Hit     byte
	Miss    byte
}

var legend = Legend{
	rune("~"),
	[]byte("0")[0],
	[]byte("X")[0],
	[]byte("*")[0],
}

/* legend is a dictionary of program logic to map elements
var legend = map[string]string{
	"terrain": "~",
	"ship":    "0",
	"hit":     "X",
	"miss":    "*",
}*/

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
		str.WriteString(fmt.Sprintf("%08b\n", l.layer[row]))
	}
	return str.String()
}

func (l *layer) coordinateToOne(c Coordinate) {
	l.layer[c.y] |= 1 << c.x
}

type board struct {
	ships layer
	shots layer
}

func CoordinatePart() {
	var ll = layer{[boardSideLength]byte{4, 8, 16, 32}}
	fmt.Println(ll)
	fmt.Printf("%T\n", ll)
	fmt.Printf("%08b\n", ll.layer)
	fmt.Printf("%s\n", ll.StringRaw())

	var c = Coordinate{5, 2}
	//c.Read()
	fmt.Printf("Internal representation: %+v\n", c)
	fmt.Println(c.String())
	//	var board = Board{}
	ll.coordinateToOne(c)
	fmt.Println(ll.StringRaw())

	fmt.Printf("%s", legend.Terrain)
	fmt.Print("%s", legend.Hit)
	fmt.Print("%s", legend.Miss)
	fmt.Print("%s", legend.Ship)

}
