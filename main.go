package main

import (
	"./lib/card"
	"fmt"
	"sort"
)

/*
 利用方法:
*/
func main() {
	var cards card.Cards
	for i := 1; i < 14; i++ {
		for j := 1; j < 5; j++ {
			card := new(card.Card)
			card.Number = i
			card.Suit.SetSuit(j)
			cards = append(cards, card)
		}
	}
	// your_hand := cards.Shuffle()[0:5]
	your_hand := cards[0:5]
	for _, c := range your_hand {
		fmt.Printf("suit:%s\tnumber:%v\n", c.Suit.SuitText, c.Number)
	}
}
