package main

import (
	"bufio"
	"deck"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type player struct {
	Score int
	Bank int
	Hand []deck.Card
}

func main() {
	fmt.Print("Number of players: ")
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	s = strings.ReplaceAll(s, "\n", "")
	n, err := strconv.Atoi(s)
	for err != nil || n == 0 {
		fmt.Println("Invalid number")
		fmt.Print("Number of players: ")
		s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		s = strings.ReplaceAll(s, "\n", "")
		n, err = strconv.Atoi(s)
	}
	players := make([]player, n+1)
	var p deck.Deck = deck.Card{
		Suit: "21",
		Name: "21",
	}
	p.Shuffle()

}

func (p *player)checkHand() {
	hand := (*p).Hand
	(*p).Score = 0
	ace := false
	var n int
	for i := 0; i < len(hand); i++ {
		switch hand[i].Name {
		case "A":
			(*p).Score += 11
			ace = true
		case "J","Q","K":
			(*p).Score += 10
		default:
			n, _ = strconv.Atoi(hand[i].Name)
			(*p).Score += n
		}
	}
	if ace && (*p).Score > 21 {
		(*p).Score -= 10
	}
}
