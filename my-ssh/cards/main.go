package main

import "fmt"


type person struct {
  firstname string
  lastname string
}


func (personPointer *person) updateName(newName string){
  (*personPointer).firstname = newName
}


func main() {

  person := person{"Alex", "Feguson"}
  fmt.Printf("Person: %+v", person)
  
  personPointer := &person
  personPointer.updateName("Hal")

  fmt.Printf("Person: %+v", person)
  

//	cards := newDeck()
//	hand, remaing := deal(cards, 5)

	/*println("---------first deck------------")

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
*/
	// println("---------clean file------------")
	// cleanFile("deck.txt")
}
