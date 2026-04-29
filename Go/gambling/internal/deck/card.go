package deck

import (
	"fmt"
	"strings"
)

type Rank int
type Suit string
type Symb rune

const (
	two   Rank = iota + 2 // start at 2
	three                 // 3
	four                  // 4
	five                  // ...
	six
	seven
	eight
	nine
	ten
	jack
	queen
	king
	ace
	joker
)

const (
	diamonds Suit = "diamonds"
	clubs    Suit = "clubs"
	hearts   Suit = "hearts"
	spades   Suit = "spades"
	wild     Suit = "whimsy" // hehe, it works \shrug
)

var names = []string{"two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "jack", "queen", "king", "ace", "joker"}
var suits = []Suit{"diamonds", "clubs", "hearts", "spades"}
var symbols = map[string]Symb{
	"two":      '2',
	"three":    '3',
	"four":     '4',
	"five":     '5',
	"six":      '6',
	"seven":    '7',
	"eight":    '8',
	"nine":     '9',
	"ten":      'X',
	"jack":     'J',
	"queen":    'Q',
	"king":     'K',
	"ace":      'A',
	"joker":    '?',
	"spades":   '♠',
	"clubs":    '♣',
	"diamonds": '♦',
	"hearts":   '♥',
}

// Represents a playing card
type Card struct {
	Rank Rank
	Suit Suit
}

func NewCard(rank Rank, suit Suit) *Card {

	card := Card{
		Rank: rank,
		Suit: suit,
	}
	return &card
}

// Checks whether these cards are identical
func (c *Card) Equals(other Card) bool {
	if c.Rank == other.Rank && c.Suit == other.Suit {
		return true
	}
	return false
}

func getSymb(keyword string) Symb {
	return Symb(symbols[keyword])
}

func getName(rank Rank) string {
	return names[rank-2]
}

// Card as text
func (c *Card) String() string {
	return fmt.Sprintf("%s of %s", getName(c.Rank), c.Suit)
}

// Card displays image of itself to terminal
func (c *Card) Display() string {
	var sb strings.Builder
	sb.WriteString("\n __ \n")
	fmt.Fprintf(&sb, "|%c |\n", getSymb(getName(c.Rank)))
	fmt.Fprintf(&sb, "| %c|\n", getSymb(string(c.Suit)))
	sb.WriteString(" ‾‾ ")
	fmt.Println(sb.String())
	return sb.String()
}

// Displays multiple cards adjacently. Allows for far
// more readable terminal output.
func DisplayMultipleCards(cards ...*Card) string {

	var sb strings.Builder
	sb.WriteString("\n\n")

	for pos := 0; pos < len(cards); pos += 5 {

		// get next five cards
		var next []*Card
		if len(cards)-pos >= 5 {
			next = cards[pos : pos+5]
		} else {
			next = cards[pos:]
		}

		// display cards adjacently
		for range next {
			sb.WriteString(" __  ")
		}
		sb.WriteRune('\n')
		for i := range next {
			fmt.Fprintf(&sb, "|%c | ", getSymb(getName(next[i].Rank)))
		}
		sb.WriteRune('\n')
		for i := range next {
			fmt.Fprintf(&sb, "| %c| ", getSymb(string(next[i].Suit)))
		}
		sb.WriteRune('\n')
		for range next {
			sb.WriteString(" ‾‾  ")
		}
		sb.WriteRune('\n')

	}

	fmt.Println(sb.String())
	return sb.String()
}
