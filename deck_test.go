package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T)  {
	d:=newDeck()
	if len(d) != 12{
		t.Errorf("Expected deck length of 16 but got %v",len(d))
	}
}
func TestSaveToDeckAndNewDeckFromFile(t *testing.T)  {
	fileName:="_deckTesting"
	os.Remove(fileName)
	deck:=newDeck()
	deck.saveToFile(fileName)
	loadedDeck:=newDeckFromFile(fileName)
	if(len(loadedDeck)!= len(deck)){
		t.Errorf("Both deck length does not match required %v,found %v",len(deck),len(loadedDeck))
	}
	os.Remove(fileName)
}