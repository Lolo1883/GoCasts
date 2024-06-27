package main

func main() {
	cards := newDeck()
	hand, remaing := deal(cards, 5)

	println("---------first deck------------")

	hand.print()

	println("---------last deck------------")
	remaing.print()

	println("---------save hand------------")
	hand.saveToFile("deck.txt")

	println("---------Reading file------------")
	fromFileDeck := readFromFile("deck.txt")
	fromFileDeck.print()

	println("---------After Shuffle------------")
	fromFileDeck.shuffle()
	fromFileDeck.print()

	// println("---------clean file------------")
	// cleanFile("deck.txt")
}
