// games package collects the code for games in Golang
package games

import (
	"crypto/rand"
	"fmt"
)

// BinaryGuessing is a game to guess random bits
func BinaryGuessing() {
	var fullExpectation, tracker uint8
	var expectation, intin int

	b := make([]byte, 1)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fullExpectation = b[0]
	fmt.Printf("FULL EXPECTATION: %08b\t", fullExpectation)
	fmt.Printf("TRACKER: %08b\n\n", tracker)

	for i := uint8(0); i < 8; i++ {
		if fullExpectation&(1<<i) != 0 {
			expectation = 1
		} else {
			expectation = 0
		}

		fmt.Printf("i= %d\t", i)
		fmt.Printf("Current Expectation = %d \t", expectation)
		fmt.Print("0 or 1? ")
		fmt.Scanln(&intin)

		if intin == expectation {
			tracker |= 1 << i
		} else {
			tracker |= 0 << i
		}
	}
	fmt.Printf("\nFinal state: %08b\n", tracker)
}
