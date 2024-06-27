package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// main deck slice
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) saveToFile(filename string) {
	f, err := os.Create(filename)
	HandleError(err)

	_, err = f.Write([]byte(d.toString()))
	HandleError(err)

	defer f.Close()
}

func cleanFile(filename string) {
	os.Remove(filename)
}

func readFromFile(filename string) deck {
	f, err := os.Open(filename)
	HandleError(err)
	defer f.Close()

	sc := bufio.NewScanner(f)
	resultDeck := deck{}

	for sc.Scan() {
		s := strings.Split(sc.Text(), ",")
		resultDeck = deck(s)
	}
	return resultDeck
}

func (d deck) shuffle() deck {
	for i, _ := range d {
		newPosition := rand.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
	return d
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
