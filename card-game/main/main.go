package main

import (
	"fmt"

	"../entities"
)

func main() {
	// Se pueden definir variables sin tipado y se asignan automaticamente
	// var card string = "Ace of Spades"

	fmt.Println(entities.GenerateDeck("Oscar", 5))
}
