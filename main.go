package main

import (
	"crypto/rand"
	"flag"
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

func main() {
	var n = flag.Int("n", 1, "rounds to play")
	var game = flag.String("game", "num", "game to play (e.g. 'num', 'bin')")
	flag.Parse()

	for i := 0; i < *n; i++ {
		switch *game {
		case "num":
			NumberGuessing()
		case "bin":
			BinaryGuessing()
		default:
			fmt.Println("Pass a flag -game with either 'num' or 'bin'")
		}
	}
}
