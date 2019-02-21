package card

import (
	"sort"
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

// ここから役チェックテスト
func TestIsOnePairSuccess(t *testing.T) {
	var cards Cards
	cards = createDeck()[0:5]
	if !cards.IsOnePair() {
		t.Fatal("error. this is not one pair")
	}
}
func TestIsTwoPairSuccess(t *testing.T) {
	var cards Cards
	cards = createDeck()[0:5]
	if !cards.IsTwoPair() {
		t.Fatal("error. this is not two pair")
	}
}
func TestIsThreeCardSuccess(t *testing.T) {
	var cards Cards
	cards = createDeck()[0:5]
	if !cards.IsThreeCard() {
		t.Fatal("error. this is not three card")
	}
}
func TestIsFourCardSuccess(t *testing.T) {
	var cards Cards
	cards = createDeck()[0:5]
	if !cards.IsFourCard() {
		t.Fatal("error. this is not four card")
	}
}
func TestIsFlushSuccess(t *testing.T) {
	var cards Cards
	cards = createDeck()
	sort.SliceStable(cards, func(i, j int) bool { return cards[i].Suit.SuitInt < cards[j].Suit.SuitInt })
	cards = cards[0:5]
	if !cards.IsFrush() {
		t.Fatal("error. this is not flush")
	}
}
func TestIsStraightSuccess(t *testing.T) {
	var cards Cards
	cards = createDeck()
	sort.SliceStable(cards, func(i, j int) bool {
		return cards[i].Suit.SuitInt < cards[j].Suit.SuitInt
	})
	cards = cards[0:5]
	if !cards.IsStraight() {
		t.Fatal("error. this is not straight")
	}

}
func TestIsStraintFlushSuccess(t *testing.T) {
	var cards Cards
	cards = createDeck()
	sort.SliceStable(cards, func(i, j int) bool {
		return cards[i].Suit.SuitInt < cards[j].Suit.SuitInt
	})
	cards = cards[0:5]
	if !cards.IsStraight() {
		t.Fatal("error. this is not straight")
	}
}

// func TestIsPokerHandSuccess(t *testing.T) {
// 	var cards Cards = createDeck()[:5]
// 	if cards.IsPokerHand() == "four card" {

// 	}
// 	t.Fatal("error")
// }
