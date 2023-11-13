package main

func main() {
	// assigns type of card
	// cards := newDeck()

	// hand, remaingDeck := deal(cards, 6)
	// hand.print()
	// remaingDeck.print()
	// cards := newDeck()

	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards.csv")
	cards := newDeckFromFile("my1_cards")
	cards.print()

}
