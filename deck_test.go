package deck

import (
	"testing"
)

func TestNew(t *testing.T) {
	got := New()
	if len(got) != 52 {
		t.Errorf("Expecting new deck of 52 cards, received deck of %d cards", len(got))
	}
	exp := Card{Suit: Spades, Value: Ace}
	if got[0] != exp {
		t.Errorf("Expecting %s as first card, received %s", exp, got[0])
	}
}

func TestWithJokers(t *testing.T) {
	tcases := map[string]struct {
		want int
	}{
		"2 Jokers":                   {want: 2},
		"0 Jokers":                   {want: 0},
		"More jokers than deck size": {want: 53},
	}
	for name, tc := range tcases {
		t.Run(name, func(t *testing.T) {
			got := New(WithJokers(tc.want))
			var count int = 0
			for _, card := range got {
				if card.Suit == Joker {
					count++
				}
			}
			if count != tc.want {
				t.Errorf("Expecting %d jokers, this deck has %d", tc.want, count)
			}
		})
	}
}

func TestNumDecks(t *testing.T) {
	tcases := map[string]struct {
		num  int
		want int
	}{
		"2 decks": {num: 2, want: 104},
		"0 decks": {num: 0, want: 0},
	}
	for name, tc := range tcases {
		t.Run(name, func(t *testing.T) {
			got := New(NumDecks(tc.num))
			if len(got) != tc.want {
				t.Errorf("Expecting %d cards, received %d cards", tc.want, len(got))
			}
		})
	}
}

func TestCleanCut(t *testing.T) {
	got := New(CleanCut)
	want := Card{Value: 1, Suit: Clubs}
	if got[0] != want {
		t.Errorf("Expected first card after cut to be %s, actual card is %s", want, got[0])
	}
}

func TestDefaultSort(t *testing.T) {
	deck := New(Shuffle)
	got := DefaultSort(deck)
	want := Card{Value: 1, Suit: Spades}
	if got[0] != want {
		t.Errorf("Expected first card after sorting to be %s, actual card is %s", want, got[0])
	}

}
func TestCustomSort(t *testing.T) {
	// Uses same sort func as DefaultSort to check custom function will be accepted.
	got := New(CustomSort(less))
	want := New()
	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("Custom sort not working, expected %s at position %d, received %s", want[i], i, got[i])
		}
	}
}

func TestFilter(t *testing.T) {
	remove := func(card Card) bool {
		return card.Value == Two || card.Suit == Spades
	}
	got := New(Filter(remove))
	for _, card := range got {
		if card.Value == Two || card.Suit == Spades {
			t.Error("Expected all twos and spades to be filtered out.")
		}
	}

}

func TestShuffle(t *testing.T) {
	got := New(Shuffle)
	_ = got
	// TODO: give Shuffle() a deterministic state to be testable.
	// var rand = NewSource(Seed(1)) ?
}
