package main

func main() {
	cards := deck{"Six of Spades", newCard()}
	cards = append(cards, "Seven of Diamonds")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
