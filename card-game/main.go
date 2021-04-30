package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	// Se pueden definir variables sin tipado y se asignan automaticamente
	// var card string = "Ace of Spades"

	fmt.Println(generateDeck())
}

//Agrega al slice una carta con un numero aleatorio y con un "palo" aleatorio
func newSuit() string {
	types := []string{"Spades", "Clubs", "Hearts", "Diamonds"}
	return (types)[rand.Intn(len(types))]
}

func generateDeck() []string {
	//Es necesario para actualizar el numero aleatorio, sino será el mismo en cada ejecución
	rand.Seed(time.Now().UnixNano())

	//Listas vacía y llena de strings
	cards := []string{}

	for i := 0; i <= 5; i++ {

		//Genero numero aleatorio para asignarselo a la slice de retorno
		cardNumber := rand.Intn(13)
		cards = append(cards, strconv.Itoa(cardNumber)+" Of "+newSuit())
	}

	return cards
}
