package main

import (
	"./lib/card"
	"fmt"
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
	your_hand := card.Shuffle(cards)[0:5]
	for _, c := range your_hand {
		fmt.Printf("suit:%s\tnumber:%v\n", c.Suit.SuitText, c.Number)
	}
}
