package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print(msg string) {
	for i, card := range d {
		fmt.Println(msg, i+1, card)
	}
}

func deal(d deck, handsize int) (deck, deck) {
	hand := d[:handsize]
	remainingDeck := d[handsize:]
	return hand, remainingDeck
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	bytes := []byte(d.toString())
	err := ioutil.WriteFile(filename, bytes, 0644)

	if err != nil {
		log.Fatal(err)
	}
	return err
}

func newDeckFromFile(filepath string) deck {
	bs, err := ioutil.ReadFile(filepath)
	if err != nil {
		// cards := newdeck()
		// cards.savetofile(deckpath)

		log.Fatal(err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffle() deck {
	now := time.Now().UnixNano()
	source := rand.NewSource(now)
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
	return d
}
