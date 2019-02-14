package main

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

/*
 利用方法:
*/
func main() {
	var cards []*Card
	for i := 1; i < 14; i++ {
		for j := 1; j < 5; j++ {
			card := new(Card)
			card.Number = i
			card.Suit.SetSuit(j)
			cards = append(cards, card)
		}
	}
	for _, c := range shuffle(cards)[0:5] {
		fmt.Printf("suit:%s\tnumber:%v\n", c.Suit.SuitText, c.Number)
	}
}
func shuffle(array []*Card) []*Card {
	rand.Seed(time.Now().UnixNano())
	var result []*Card
	for i := 0; i < len(array); i++ {
		result = insert(result, rand.Intn(len(array)), array[i])
	}
	return result
}
func insert(slice []*Card, pos int, value *Card) []*Card {
	if len(slice) <= pos {
		return append(slice, value)
	}
	array := append(slice[:pos+1], slice[pos:]...)
	array[pos] = value
	return array
}
