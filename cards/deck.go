// package main

// import "fmt"

// type deck []string

// func (d deck) print(msg string) {
// 	for i, card := range d {
// 		fmt.Println(msg, i, card)
// 	}
// }

// func newDeck() deck {
// 	cards := deck{}

// 	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
// 	cardValues := []string{"Ace", "Two", "Three", "Four"}

// 	for _, suit := range cardSuits {
// 		for _, value := range cardValues {
// 			cards = append(cards, value+" of "+suit)
// 		}
// 	}
// 	return cards
// }
