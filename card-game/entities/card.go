package entities

import (
	"math/rand"
	"time"
)

type Card struct {
	Number int
	Suit   string
}

func NewCard() Card {

	//Actualiza el numero aleatorio, sino será el mismo en cada ejecución
	rand.Seed(time.Now().UnixNano())

	return Card{
		Number: rand.Intn(13), Suit: newSuit(),
	}
}

//Agrega al slice una carta con un numero aleatorio y con un "palo" aleatorio
func newSuit() string {

	types := []string{"Spades", "Clubs", "Hearts", "Diamonds"}

	return (types)[rand.Intn(len(types))]
}
