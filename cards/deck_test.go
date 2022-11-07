package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expect deck length of 16 but got %v", len(d))
	}
}

func TestSaveToFileNewDeckFromFile(t *testing.T) {
	filename := "_testfile"
	os.Remove(filename)
	d := newDeck()
	d.saveToFile(filename)
	d = newDeckFromFile(filename)

	if len(d) != 16 {
		t.Errorf("Expect deck length of 16 but got %v", len(d))
	}

	os.Remove(filename)
}
