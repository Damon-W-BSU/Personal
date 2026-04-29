package deck

import (
	"testing"
)

func TestNewCard(t *testing.T) {
	card := NewCard(two, spades)
	t.Logf("%v", card)
	t.Logf("\n%v\n", card.Display())
	if card.Rank != 2 || card.Suit != "spades" {
		t.Fail()
	}
}

func TestDisplayMultiple(t *testing.T) {

	card1 := NewCard(two, diamonds)
	card2 := NewCard(king, diamonds)

	t.Log(DisplayMultipleCards(card1))
	t.Log(DisplayMultipleCards(card1, card2))

}
