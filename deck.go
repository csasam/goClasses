package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of string
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{
		"Spades", "Diamonds", "Hearts", "Clubs",
	}

	cardValues := []string{
		"Ace", "Two", "Three", "Four",
	}

	// replace variable with underscore tells go
	// its a variable but we arent going to use it
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards

}

func (d deck) print() {
	// this is a reciever that allows anything of type deck access to the print method
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	// returns error b/c writefile returns error if there is error

	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)

}
