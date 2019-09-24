package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"sort"
	"time"
)

func main() {
	dealer := InitDealer()
	//fmt.Println(len(dealer.Deck.Deck))
	//dealer.Cut(26)
	//dealer.PrintDeck()
	//fmt.Printf("dealt card: %v\n", dealer.DealCard())
	//fmt.Printf("dealt card: %v\n", dealer.DealCard())
	//fmt.Printf("Dealt Cards: %v\n", dealer.DealtCards )
	//fmt.Println(len(dealer.Deck.Deck))
	//dealer.Shuffle()
	//dealer.PrintDeck()
	//fmt.Printf("dealt card: %v\n", dealer.DealCard())
	//fmt.Printf("dealt card: %v\n", dealer.DealCard())
	//
	//dealer.PrintDeck()
	//dealer = RebuildDeck()
	//dealer.PrintDeck()
	//fmt.Printf("dealt card: %v\n", dealer.DealCard())
	//fmt.Printf("dealt card: %v\n", dealer.DealCard())
	//fmt.Printf("Dealt Cards: %v\n", dealer.DealtCards )
	//
	//dealer.Discard(1)
	//fmt.Printf("Discard Deck: %v\n", dealer.DiscardDeck)
	//
	//dealer = InitDealer()
	//fmt.Printf("dealt card: %v\n", dealer.DealCard())
	//fmt.Printf("dealt card: %v\n", dealer.DealCard())
	//dealer.PrintDeck()
	//dealer.Shuffle()
	//dealer.PrintDeck()
	//dealer.Sort()
	//dealer.PrintDeck()

	router := mux.NewRouter()

	router.HandleFunc("/DealCard", func(w http.ResponseWriter, r *http.Request) {
		dealer.PrintDeck()
		card := dealer.DealCard()
		dealer.PrintDeck()
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(card.String())
	}).Methods("GET")

	http.ListenAndServe("localhost:3001", router)
}

type Dealer struct {
	Deck Deck
	DealtCards []Card
	DiscardDeck []Card
}

type Suit int

const (
	Spades Suit = iota
	Hearts
	Clubs
	Diamonds
)

func (s Suit) String() string {
	switch s {
	case Spades:
		return "Spades"
	case Hearts:
		return "Hearts"
	case Clubs:
		return "Clubs"
	case Diamonds:
		return "Diamonds"
	default:
		return fmt.Sprintf("%d", int(s))
	}
}

type Face int

func (f Face) String() string {
	switch f {
	case Duece:
		return "2"
	case Three:
		return "3"
	case Four:
		return "4"
	case Five:
		return "5"
	case Six:
		return "6"
	case Seven:
		return "7"
	case Eight:
		return "8"
	case Nine:
		return "9"
	case Ten:
		return "10"
	case Jack:
		return "Jack"
	case Queen:
		return "Queen"
	case King:
		return "King"
	case Ace:
		return "Ace"
	default:
		return fmt.Sprintf("%d", int(f))
	}
}

const (
	Duece Face = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

type Deck struct {
	Deck []Card
}

type Card struct {
	Face
	Suit
}

func (c Card) String() string {
	return fmt.Sprintf( "%s of %s", c.Face, c.Suit)
}

func InitDealer() Dealer {
	deck := initDeck()
	return Dealer{
		Deck: deck,
		DealtCards: []Card{},
		DiscardDeck: []Card{},
	}
}

func initDeck() Deck {
	deck := make([]Card, 0)
	suits := []Suit{Spades, Hearts, Clubs, Diamonds}
	faces := []Face{Duece, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}
	for _, suit := range suits {
		//fmt.Println(suit)
		for _, face := range faces {
			//fmt.Println(face)
			deck = append(deck, Card{face, suit})
		}
	}
	return Deck{
		Deck: deck,
	}
}

func (d *Dealer) PrintDeck() {
	fmt.Printf("Deck: %v\n", d.Deck.Deck)
}

// take in a slice or queue and pop the top value
func (d *Dealer) DealCard() Card {
	// Deal from the top of the deck
	card := d.Deck.Deck[0] // pop from the front of the slice
	d.Deck.Deck = d.Deck.Deck[1:] // remove card from slice
	d.DealtCards = append(d.DealtCards, card)

	return card
}

// Randomizing function taken from https://www.calhoun.io/how-to-shuffle-arrays-and-slices-in-go/ 2nd solution from the bottom
func (d *Dealer) Shuffle() {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	// We start at the end of the slice, inserting our random
	// values one at a time.
	for n := len(d.Deck.Deck); n > 0; n-- {
		randIndex := r.Intn(n)
		// We swap the value at index n-1 and the random index
		// to move our randomly chosen value to the end of the
		// slice, and to move the value that was at n-1 into our
		// unshuffled portion of the slice.
		d.Deck.Deck[n-1], d.Deck.Deck[randIndex] = d.Deck.Deck[randIndex], d.Deck.Deck[n-1]
	}
}

func (d *Dealer) Cut(pos int) {
	// check that position is less than size of deck
	if pos > len(d.Deck.Deck) {
		fmt.Printf("please choose a position in the deck.\nThe deck shrinks when DealCard is called.\n")
		return
	}

	// swap at specified position
	d.Deck.Deck = append(d.Deck.Deck[pos:], d.Deck.Deck[:pos]...)
}

func RebuildDeck() Dealer {
	return InitDealer()
}

func (d *Dealer) Discard(pos int) {
	// edge case size is 0
	if len(d.DealtCards) == 0 {
		fmt.Println("Must deal a card before it can be discarded")
		return
	}

	d.DiscardDeck = append(d.DealtCards[:pos], d.DealtCards[pos+1:]...)
}

// Sort info obatained from stackoverflow question: https://stackoverflow.com/questions/36122668/how-to-sort-struct-with-multiple-sort-parameters
func (d *Dealer) Sort() {
	sort.Slice(d.Deck.Deck, func(i, j int) bool {
		if d.Deck.Deck[i].Suit < d.Deck.Deck[j].Suit {
			return true
		}
		if d.Deck.Deck[i].Suit > d.Deck.Deck[j].Suit {
			return false
		}
		return d.Deck.Deck[i].Face < d.Deck.Deck[j].Face
	})
}