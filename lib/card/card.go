package card

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Card struct {
	Number int
	Suit   Suit
}
type Suit struct {
	SuitText string
	SuitInt  int
}

func (s *Suit) SetSuit(n int) {
	s.SuitInt = n
	switch n {
	case 1:
		s.SuitText = "spade"
	case 2:
		s.SuitText = "club"
	case 3:
		s.SuitText = "diamond"
	case 4:
		s.SuitText = "heart"
	default:
		fmt.Println("not found")
	}
}

func (cards Cards) Len() int {
	return len(cards)
}

func (cards Cards) Swap(i, j int) {
	if i == j {
		return
	}
	*cards[i], *cards[j] = *cards[j], *cards[i]
}

func (cards Cards) Less(i, j int) bool {
	return cards[i].Number < cards[j].Number
}

type Cards []*Card

func (cards Cards) Shuffle() Cards {
	var result Cards
	result = make([]*Card, len(cards))
	copy(result, cards)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < cards.Len(); i++ {
		rand.Seed(time.Now().UnixNano() + int64(i))
		result.Swap(i, rand.Intn(cards.Len()))
	}
	return result
}
func (cards Cards) Insert(pos int, value *Card) Cards {
	if len(cards) <= pos {
		return append(cards, value)
	}
	array := append(cards[:pos+1], cards[pos:]...)
	array[pos] = value
	return array
}

func (cards Cards) IsOnePair() bool {
	sort.SliceStable(cards, func(i, j int) bool { return cards[i].Suit.SuitInt < cards[j].Suit.SuitInt })
	for i, card_i := range cards {
		for _, card_j := range cards[i+1:] {
			if card_i.Number == card_j.Number {
				return true
			}
		}
	}
	return false
}
func (cards Cards) IsTwoPair() bool {
	sort.SliceStable(cards, func(i, j int) bool { return cards[i].Suit.SuitInt < cards[j].Suit.SuitInt })
	for i, card_i := range cards {
		for j, card_j := range cards[i+1:] {
			if card_i.Number == card_j.Number {
				removed_cards := append(cards[:i], cards[i+1:]...)
				removed_cards = append(removed_cards[:j], removed_cards[j+1:]...)
				return removed_cards.IsOnePair()
			}
		}
	}
	return false
}
func (cards Cards) IsThreeCard() bool {
	sort.SliceStable(cards, func(i, j int) bool { return cards[i].Suit.SuitInt < cards[j].Suit.SuitInt })
	for i, card_i := range cards {
		cnt := 1
		for _, card_j := range cards[i+1:] {
			if card_i.Number == card_j.Number {
				cnt++
			}
			if cnt >= 3 {
				return true
			}
		}
	}
	return false
}

// TODO: 13,1,2,3,4みたいな奴は現時点では非サポートなので後ほどサポートするようにする。
func (cards Cards) IsStraight() bool {
	sort.SliceStable(cards, func(i, j int) bool { return cards[i].Suit.SuitInt < cards[j].Suit.SuitInt })
	for i := 0; i < len(cards)-1; i++ {
		if cards[i].Number+1 != cards[i+1].Number {
			return false
		}
	}
	return true
}
func (cards Cards) IsFrush() bool {
	firstCard := cards[0]
	for _, card := range cards[1:] {
		if firstCard.Suit.SuitInt != card.Suit.SuitInt {
			return false
		}
	}
	return true
}
func (cards Cards) IsFullHouse() bool {
	return !cards.IsFourCard() && cards.IsTwoPair() && cards.IsThreeCard()
}
func (cards Cards) IsFourCard() bool {
	sort.SliceStable(cards, func(i, j int) bool { return cards[i].Suit.SuitInt < cards[j].Suit.SuitInt })
	for i, card_i := range cards {
		cnt := 1
		for _, card_j := range cards[i+1:] {
			if card_i.Number == card_j.Number {
				cnt++
			}
			if cnt >= 4 {
				return true
			}
		}
	}
	return false
}
func (cards Cards) IsStraightFlush() bool {
	return cards.IsFrush() && cards.IsStraight()
}

// 役文字列を返す。
func (cards Cards) IsPokerHand() string {
	if cards.IsStraightFlush() {
		return "straight flush"
	} else if cards.IsFourCard() {
		return "four card"
	} else if cards.IsFullHouse() {
		return "fullhouse"
	} else if cards.IsFrush() {
		return "flush"
	} else if cards.IsStraight() {
		return "straihght"
	} else if cards.IsThreeCard() {
		return "three card"
	} else if cards.IsTwoPair() {
		return "two pair"
	} else if cards.IsOnePair() {
		return "one pair"
	}
	return "none"
}
