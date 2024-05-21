package main

import (
	"fmt"
	ph "github.com/forrest321/pokerHands"
)

func main() {
	run()
}

func run() {
	var hands []ph.Hand
	deck := ph.NewDeck()
	deck.Shuffle()
	for i := 0; i < 5; i++ {
		hand := ph.Hand{}
		hand.Cards = deck.Deal(0, 5)
		hand.Evaluate()
		fmt.Println(hand.String())
		hands = append(hands, hand)
	}
	winners := ph.FindWinners(hands)
	fmt.Println("winners")
	for _, w := range winners {
		fmt.Println(w.String())
	}
}
