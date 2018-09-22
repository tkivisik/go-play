package games

import (
	"crypto/rand"
	"fmt"
)

// NumberGuessing is a game to guess random numbers.
func NumberGuessing() {
	var intin int

	key := [100]byte{}
	_, err := rand.Read(key[:])
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(key); i++ {
		if key[i] <= 100 {
			fmt.Println("Guess the number in range [1, 100]")
			for {
				fmt.Scan(&intin)
				if intin > int(key[i]) {
					fmt.Println("\tLower")
				} else if intin < int(key[i]) {
					fmt.Println("\tHigher")
				} else {
					fmt.Println("\tCorrect!")
					break
				}
			}
			break
		}
	}
}
