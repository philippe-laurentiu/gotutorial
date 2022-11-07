package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	a := []string{"Spades", "Diamonts", "Hearts", "Cross"}
	b := []string{"Ace", "One", "Two", "Three"}

	for _, valuea := range a {
		for _, valueb := range b {
			cards = append(cards, valuea+" "+valueb)
		}
	}
	return cards
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffel() {
	rs := rand.NewSource(time.Now().UnixNano())
	random := rand.New(rs)

	for i := range d {
		r := random.Intn(len(d) -1)
		d[i], d[r] = d[r], d[i]
	}
}