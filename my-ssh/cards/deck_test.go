package main

import (
   "testing"
   "os"
)


func TestNewDeck(t *testing.T) {
   d := newDeck()

   if len(d) != 16 {
       t.Errorf("Expected deck length of 16 but got %d", len(d))
   }
}


func TestFirstInADeck(t *testing.T) {
  d := newDeck()
  if len(d) == 16 && d[0] != "Ace of Spades" {
    t.Errorf("Expected first card in a deck Ace of Spades but got %s", d[0])
  }
}

func TestLastInADeck(t *testing.T) {
  d := newDeck()
  if len(d) == 16 && d[len(d) - 1] != "Four of Clubs" {
    t.Errorf("Expected first card in a deck Ace of Spades but got %s", d[len(d) - 1])
  }
}

func TestSaveAndLoadDeckFromFiles(t *testing.T) {
  tmpDeck := newDeck()
  tmpDeck.saveToFile("testDeck.txt")
  fromFileDeck := readFromFile("testDeck.txt")
  
  if len(fromFileDeck) == 0 {
    t.Error("Expected a deck loaded from file but found nothing")
  }

  teardown()
}


func teardown(){
  os.Remove("testDeck.txt")
}
