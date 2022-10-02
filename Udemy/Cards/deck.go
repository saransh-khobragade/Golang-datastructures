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

func getDeck() deck {

	cards := deck{}

	temp1 := []string{"A", "B", "C"}
	temp2 := []string{"a", "b", "c"}

	for _, a := range temp1 {
		for _, b := range temp2 {
			cards = append(cards, a+b)
		}
	}

	return cards
}

func (d deck) print() { //any variable with type deck now can access this function
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

func newDeckFormFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {

	for index := range d {
		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)
		i := r.Intn(len(d) - 1)
		d[index], d[i] = d[i], d[index]
	}
}
