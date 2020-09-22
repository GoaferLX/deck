//go:generate stringer -type=Suit,Value
package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Suit int

const (
	Spades Suit = iota
	Diamonds
	Clubs
	Hearts
	Joker
)

var suits = [4]Suit{Spades, Diamonds, Clubs, Hearts}

type Value int

const (
	Ace Value = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

type Card struct {
	Value
	Suit
}

// String to satisfy Stringer interface.
// Returns card as a human readable string.
func (c Card) String() string {
	if c.Suit == Joker {
		return fmt.Sprint("Joker")
	}
	return fmt.Sprintf("%s of %s", c.Value.String(), c.Suit.String())
}

// New generates a new deck of Cards.
// 52 cards standard - no Jokers.
// Accepts deckOpts to provide a more defined deck.
// type deckOpts func([]Card) []Card.
func New(opts ...func([]Card) []Card) []Card {
	var deck []Card
	for _, suit := range suits {
		var value Value
		for value = 1; value <= 13; value++ {
			deck = append(deck, Card{Value: value, Suit: suit})
		}
	}
	for _, opt := range opts {
		deck = opt(deck)
	}
	return deck
}

// Optional functions for creating a new deck.

// Filter allows user to provide custom filter functions
// to filter specific cards i.e. no Twos.
func Filter(fn func(Card) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		var ret []Card
		for _, card := range cards {
			if !fn(card) {
				ret = append(ret, card)
			}
		}
		return ret
	}
}

// NumDeck creates a larger deck consisting of multiple decks.
// i.e. Casino Blackjack uses 6 decks for standard play.
func NumDecks(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		var ret []Card
		for i := 0; i < n; i++ {
			for _, card := range cards {
				ret = append(ret, card)
			}
		}
		return ret
	}
}

// Add arbitrary number of jokers to the deck
func WithJokers(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{Suit: Joker})
		}
		return cards
	}
}

//	Shuffle mixes the cards and returns the original slice in a random order.
func Shuffle(cards []Card) []Card {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
	return cards
}

// Cut cuts the deck and completes the cut at a random point.
func Cut(cards []Card) []Card {
	n := rand.Intn(len(cards))
	cards = append(cards[n:], cards[:n]...)
	return cards
}

// CleanCut cuts the deck in the middle and completes the cut.
func CleanCut(cards []Card) []Card {
	n := len(cards) / 2
	cards = append(cards[n:], cards[:n]...)
	return cards
}

// DefaultSort sorts cards into new deck order.
// Ace -> King Spades, Ace-> King Diamonds, Ace -> King Clubs, Ace -> King Hearts.
func DefaultSort(cards []Card) []Card {
	sort.SliceStable(cards, less(cards))
	return cards
}

//  less - returns a less function for sort.Interface.
func less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return (13*int(cards[i].Suit))+int(cards[i].Value) < (13*int(cards[j].Suit) + int(cards[j].Value))
	}
}

// CustomSort allows the user to implement a custom sorting option.
func CustomSort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.SliceStable(cards, less(cards))
		return cards
	}
}
