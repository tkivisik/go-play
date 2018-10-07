package coordinate

import (
	"fmt"
	"strconv"
	"strings"
)

//const letters = [26]byte{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
const (
	letters         string = "abcdefghijklmnopqrstuvwxyz"
	boardSideLength uint8  = 4
)

// coordinate expresses a location on a map as a
type Coordinate struct {
	x byte
	y byte
}

func (c *Coordinate) String() string {
	return fmt.Sprintf("%c%d", byte(letters[c.x]), c.y)
}

func (c *Coordinate) Read() {
	var s string
	for i := 0; i < 100; i++ {
		fmt.Scanln(&s)
		if strings.Contains(letters[boardSideLength:], string(s[0])) {
			fmt.Println("Please be careful")
			continue
		}
		y, err := strconv.Atoi(s[1:])
		if err != nil {
			fmt.Println("Please be careful")
			continue
		}
		if y <= 0 || uint8(y) > boardSideLength {
			fmt.Println("Please be careful")
			continue
		}
		c.x = byte(strings.IndexRune(letters, rune(s[0])))
		c.y = byte(y)
		c.y--
		break
	}
}

func CoordinatePart() {
	var c Coordinate = Coordinate{2, 2}
	c.Read()
	fmt.Println(c)
	fmt.Println(c.String())
}
