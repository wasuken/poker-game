package card

import (
	"fmt"
	"math/rand"
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
	cards[i], cards[j] = cards[j], cards[i]
}

func (cards Cards) Less(i, j int) bool {
	return cards[i].Number < cards[j].Number
}

type Cards []*Card

func (cards Cards) Shuffle() Cards {
	rand.Seed(time.Now().UnixNano())
	var result Cards
	for i := 0; i < cards.Len(); i++ {
		result = result.Insert(rand.Intn(cards.Len()), cards[i])
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
