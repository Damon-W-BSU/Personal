package deck

type Deck struct {
	cards  []*Card
	jokers bool
}

// Creates a standard 52 card deck
func NewDeck() *Deck {

	deck := Deck{
		cards: []*Card{},
	}

	// add standard 52 cards
	for _, suit := range suits {
		for rank := range 13 {
			deck.cards = append(deck.cards, NewCard(Rank(rank+2), suit))
		}
	}

	return &deck
}

// Returns a copy of the deck sorted by suit
func (d *Deck) SortBySuit() *Deck {
	return nil
}

// Returns a copy of the deck sorted by rank
func (d *Deck) SortByRank() *Deck {
	return nil
}

// TODO
// func (d *Deck) DeleteCard(card Card) *Card {
// 	for _, card := range d.cards {
// 		if card.Equals(*card)
// 	}
// }

// TODO
// // Returns a copy of the deck with order randomized
// func (d *Deck) Shuffle() *Deck {

// 	shuffled := []*Card{}
// 	items := len(d.cards)

// 	for i := range items {
// 		shuffled = append(shuffled, d.cards[rand.Int()%len(d.cards)])
// 	}

// }

// Displays all cards in the deck to terminal
func (d *Deck) Display() string {
	return DisplayMultipleCards(d.cards...)
}
