package card

import (
	"testing"
)

func createDeck() Cards {
	var cards Cards
	for i := 1; i < 14; i++ {
		for j := 1; j < 5; j++ {
			card := new(Card)
			card.Number = i
			card.Suit.SetSuit(j)
			cards = append(cards, card)
		}
	}
	return cards
}

func TestInsertSuccess(t *testing.T) {
	cards := createDeck()
	card := new(Card)
	card.Number = 10
	card.Suit.SetSuit(3)
	cards.Insert(0, card)
	if cards[0].Number != 10 || card.Suit.SuitText != "diamond" {
		t.Fatal("error")
	}
}
func TestShuffleSuccess(t *testing.T) {
	cards := createDeck()
	// 流石にこれくらいやれば大丈夫やろ多分、
	if cards.Shuffle()[0] == cards.Shuffle()[0] && cards.Shuffle()[0] == cards.Shuffle()[0] {
		t.Fatal("error")
	}
}

func TestSwapSuccess(t *testing.T) {
	cards := createDeck()
	swaped_cards := make([]*Card, cap(cards))
	copy(swaped_cards, cards)
	cards.Swap(0, 1)
	if swaped_cards[0] != cards[1] || swaped_cards[1] != cards[0] {
		t.Fatal("error")
	}
}
func TestLenSuccess(t *testing.T) {
	cards := createDeck()
	if cards.Len() != 52 {
		t.Fatal("error")
	}
}
func TestLessSuccess(t *testing.T) {
	cards := createDeck()
	if cards.Less(0, 1) {
		t.Fatal("error")
	}
}
