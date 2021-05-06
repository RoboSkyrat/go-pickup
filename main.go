package main

import (
	"bufio"
	"fmt"
	"github.com/RoboSkyrat/pickup/card"
	"os"
	"strings"
	"time"
)

func main() {

	var cards map[int]bool = map[int]bool{}
	start := time.Now()
	fmt.Print("There are 52 cards on the floor.\n> ")


	// Scanner reads stdin
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()


		// Split the input into words
		args := strings.Split(input, " ")

		// Determines where in the string is the start of the card
		index := 0

		if strings.ToLower(args[0]) == "pickup" {
			index++ // Skip the "pickup" part of the list so we know which index in the array is the face
		}

		suit := card.GetSuit(args[len(args)-1])
		face := card.GetFace(args[index])

		value := int(suit) + int(face)

		if cards[value] || value == 0{
			fmt.Println("That card isn't on the floor. 😔")
		} else {
			cards[value] = true
		}

		if len(cards) == 52 {
			fmt.Println("You picked up all the cards. 😃")
			fmt.Printf("It took you %s to pick them up.", time.Since(start))
			os.Exit(0)
		} else {
			fmt.Printf("There are %d cards on the floor.\n> ", 52-len(cards))
		}
	}

	if err := scanner.Err(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
