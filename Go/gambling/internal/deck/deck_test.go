package deck

import (
	"fmt"
	"strings"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := NewDeck()

	var sb strings.Builder
	for _, card := range deck.cards {
		fmt.Fprintf(&sb, "%s, ", card.Display())
	}

}

func TestDisplay(t *testing.T) {
	deck := NewDeck()
	deck.Display()
}
