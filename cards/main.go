package main

import (
	"fmt"
)

func main() {
	deckPath := "output/deckOfCards.txt"

	deckOfCards := newDeckFromFile(deckPath)

	shuffledCards := deckOfCards.shuffle()

	fmt.Println(shuffledCards)
}
