package main

import (
	"bufio"
	"deck"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type player struct {
	Hand   deck.Deck
	Points int
	Status string
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Enter number of players.")
	fmt.Print("-> ")
	n := returnValidNum(math.MaxInt32)
	players := make([]player, n)
	fmt.Println("Enter number of decks.")
	fmt.Print("-> ")
	n = returnValidNum(8)
	bjDeck := deck.New(n)
	bjDeck.Shuffle()
	var dealer player
	dealer.Hand.TakeCard(&bjDeck, 1)
	dealer.getPoints()
	dealer.Status = "playing"
	for i := 0; i < len(players); i++ {
		players[i].Hand.TakeCard(&bjDeck, 2)
		players[i].getPoints()
		players[i].Status = "playing"
	}
	fmt.Printf("\nDealer's hand: %s\n", dealer.getHand())
	fmt.Printf("         points: %d\n", dealer.Points)
	for i := 0; i < len(players); i++ {
		fmt.Printf("\n%d%s player's hand: %s\n", i+1, cardToOrdNum(i+1), players[i].getHand())
		fmt.Printf("             points: %d\n", players[i].Points)
	}
	for i := 0; i < len(players); i++ {
		if players[i].hasBlackJack() {
			fmt.Printf("\n%d%s player has Blackjack\n", i+1, cardToOrdNum(i+1))
			players[i].Status = "won"
		}
	}
	for i := 0; i < len(players); i++ {
		if players[i].Status == "playing" {
			fmt.Printf("\n%d%s player's turn.\n", i+1, cardToOrdNum(i+1))
			loop := true
			for loop {
				fmt.Println("Choose action: ")
				fmt.Println("1. Hit.")
				fmt.Println("2. Stand.")
				fmt.Print("-> ")
				n = returnValidNum(2)
				switch n {
				case 1:
					players[i].Hand.TakeCard(&bjDeck, 1)
					players[i].getPoints()
				case 2:
					loop = false
				}
				fmt.Printf("Hand: %s\n", players[i].getHand())
				fmt.Printf("Points: %d\n", players[i].Points)
				if players[i].Points > 21 {
					players[i].Status = "lost"
					loop = false
				}
			}
		}
	}
	for dealer.Points < 17 {
		dealer.Hand.TakeCard(&bjDeck, 1)
		dealer.getPoints()
	}
	fmt.Printf("\nDealer's hand: %s\n", dealer.getHand())
	fmt.Printf("         points: %d\n", dealer.Points)
	if dealer.Points == 21 {
		if dealer.hasBlackJack() {
			fmt.Println("Dealer has Blackjack.")
		}
		dealer.Status = "won"
	}
	if dealer.Points > 21 {
		dealer.Status = "lost"
	}
	for i := 0; i < len(players); i++ {
		switch players[i].Status {
		case "lost":
			switch dealer.Status {
			case "lost":
				fmt.Printf("%d%s player: draw.\n", i+1, cardToOrdNum(i+1))
			default:
				fmt.Printf("%d%s player: lost.\n", i+1, cardToOrdNum(i+1))
			}
		case "playing":
			switch dealer.Status {
			case "lost":
				fmt.Printf("%d%s player: won.\n", i+1, cardToOrdNum(i+1))
			case "playing":
				if players[i].Points > dealer.Points {
					fmt.Printf("%d%s player: won.\n", i+1, cardToOrdNum(i+1))
				}
				if players[i].Points == dealer.Points {
					fmt.Printf("%d%s player: draw.\n", i+1, cardToOrdNum(i+1))
				}
				if players[i].Points < dealer.Points {
					fmt.Printf("%d%s player: lost.\n", i+1, cardToOrdNum(i+1))
				}
			default:
				fmt.Printf("%d%s player: lost.\n", i+1, cardToOrdNum(i+1))
			}
		case "won":
			switch dealer.Status {
			case "won":
				fmt.Printf("%d%s player: draw.\n", i+1, cardToOrdNum(i+1))
			default:
				fmt.Printf("%d%s player: won.\n", i+1, cardToOrdNum(i+1))
			}
		}
	}
}

func (p *player) getPoints() {
	hand := (*p).Hand
	(*p).Points = 0
	ace := false
	var n int
	for i := 0; i < len(hand); i++ {
		switch hand[i].Name {
		case "A":
			if ace || (*p).Points+11 > 21 {
				(*p).Points += 1
			} else {
				(*p).Points += 11
				ace = true
			}
		case "J", "Q", "K":
			(*p).Points += 10
		default:
			n, _ = strconv.Atoi(hand[i].Name)
			(*p).Points += n
		}
	}
}

func (p player) getHand() string {
	s := p.Hand[0].Name
	for i := 1; i < len(p.Hand); i++ {
		s = s + ", " + p.Hand[i].Name
	}
	return s
}

func cardToOrdNum(c int) string {
	var o string
	if c > 10 && c < 20 {
		o = "th"
	} else {
		switch c % 10 {
		case 1:
			o = "st"
		case 2:
			o = "nd"
		case 3:
			o = "rd"
		default:
			o = "th"
		}
	}
	return o
}

func (p player) hasBlackJack() bool {
	if p.Points == 21 && len(p.Hand) == 2 {
		return true
	}
	return false
}

func returnValidNum(max int) int {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	s = strings.ReplaceAll(s, "\n", "")
	n, _ := strconv.Atoi(s)
	for n < 1 || n > max {
		fmt.Println("Invalid number")
		fmt.Print("-> ")
		s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		s = strings.ReplaceAll(s, "\n", "")
		n, _ = strconv.Atoi(s)
	}
	return n
}
