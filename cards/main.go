package main

func main() {
	cards := deck{"Ace of diamonds", newCard()}
	cards = append(cards, "Six of spades")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
