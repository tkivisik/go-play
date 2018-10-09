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

type Board struct {
}

func CoordinatePart() {
	/*var c Coordinate = Coordinate{2, 2}
	c.Read()
	fmt.Printf("Internal representation: %+v\n", c)
	fmt.Println(c.String())*/
	//	var board = Board{}

}
