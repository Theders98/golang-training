package entities

type Deck struct {
	Name  string
	Cards []Card
}

func GenerateDeck(name string, quantity int) Deck {

	cards := []Card{}

	for i := 0; i < quantity; i++ {
		cards = append(cards, NewCard())
	}

	return Deck{Name: name, Cards: cards}
}
