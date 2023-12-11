package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	hand := make([]Hand, 0)
	for sc.Scan() {
		hand = append(hand, ParseHand(sc.Text()))
	}

	sort.Slice(hand, func(i, j int) bool {
		if hand[i].set.handType == hand[j].set.handType {
			for k := 0; k < 5; k++ {
				if hand[i].set.hand[k] == hand[j].set.hand[k] {
					continue
				}

				return cardToIndex[hand[i].set.hand[k]] < cardToIndex[hand[j].set.hand[k]]
			}
		}

		return hand[i].set.handType > hand[j].set.handType
		})

	for _, h := range(hand) {
		fmt.Println(h)
	}
	// fmt.Println(hand)

	out := 0
	for i, h := range(hand) {
		out += (i+1) * h.score
	}
	fmt.Println(out)

}

type Hand struct {
	set Set
	score int
}

type Set struct {
	handType int
	hand string

}

const (
	FiveOfAKind = iota
	FourOfAKind
	FullHouse
	ThreeOfAKind
	TwoPair
	OnePair
	HighCard
)

var cardToIndex map[byte]int = map[byte]int{
	'J': 0,
	'2': 1,
	'3': 2,
	'4': 3,
	'5': 4,
	'6': 5,
	'7': 6,
	'8': 7,
	'9': 8,
	'T': 9,
	'Q': 10,
	'K': 11,
	'A': 12,
}

func ParseHand(s string) Hand {
	split := strings.Split(s, " ")
	if len(split) != 2 {
		panic("invalid hand")
	}

	score, _ := strconv.Atoi(split[1])

	return Hand{
		set: ParseSet(split[0]),
		score: score,
	}
}

func ParseSet(h string) Set {
	if len(h) != 5 {
		panic("invalid hand")
	}
	numOfCards := len(cardToIndex)
	count := make([]int, numOfCards)

	for i := range(h) {
		count[cardToIndex[h[i]]]++
	}

	numberOfJokers := count[cardToIndex['J']]
	if numberOfJokers == 0 {
		return ParseSetNoJoker(h)
	}

	// if 5 joker
	if numberOfJokers == 5 {
		return Set{
			handType: FiveOfAKind,
			hand: h,
		}
	}

	if numberOfJokers == 4 {
		return Set{
			handType: FiveOfAKind,
			hand: h,
		}
	}

	if numberOfJokers == 3 {
		// if there is a pair make 5 of a kind
		for i := range(h) {
			if count[cardToIndex[h[i]]] == 2 {
				return Set{
					handType: FiveOfAKind,
					hand: h,
				}
			}
		}
		// if no pair make 4 of a kind
		return Set{
			handType: FourOfAKind,
			hand: h,
		}
	}

	if numberOfJokers == 2 {
		// if there is a three of a kind make 5 of a kind
		for i := range(h) {
			if count[cardToIndex[h[i]]] == 3 {
				return Set{
					handType: FiveOfAKind,
					hand: h,
				}
			}
		}

		// if there is another pair make four of a kind
		for i := range(h) {
			if count[cardToIndex[h[i]]] == 2 && h[i] != 'J' {
				return Set{
					handType: FourOfAKind,
					hand: h,
				}
			}
		}

		// if not, make 3 of a king
		return Set{
			handType: ThreeOfAKind,
			hand: h,
		}
	}

	if numberOfJokers == 1 {
		// if there is a three of a kind make 5 of a kind
		for i := range(h) {
			if count[cardToIndex[h[i]]] == 4 {
				return Set{
					handType: FiveOfAKind,
					hand: h,
				}
			}
		}

		// if there is another pair make four of a kind
		for i := range(h) {
			if count[cardToIndex[h[i]]] == 3 {
				return Set{
					handType: FourOfAKind,
					hand: h,
				}
			}
		}
		// if there are 2 pairs make 3 of a kind
		numOfPairs := 0
		for i := 0; i < numOfCards; i++ {
			if count[i] == 2{
				numOfPairs++
			}
		}

		if numOfPairs == 2 {
			return Set{
				handType: FullHouse,
				hand: h,
			}
		}

		if numOfPairs == 1 {
			return Set{
				handType: ThreeOfAKind,
				hand: h,
			}
		}

		// if not then pair
		return Set{
			handType: OnePair,
			hand: h,
		}
	}


	return Set{
		handType: HighCard,
		hand: h,
	}
}

func ParseSetNoJoker(h string) Set {
	if len(h) != 5 {
		panic("invalid hand")
	}
	numOfCards := len(cardToIndex)
	count := make([]int, numOfCards)

	for i := range(h) {
		count[cardToIndex[h[i]]]++
	}

	numberOfJokers := count[cardToIndex['J']]
	if numberOfJokers != 0 {
		panic("Joker found")
	}

	// check if five of a kind
	for i := 0; i < numOfCards; i++ {
		if count[i] == 5 {
			return Set{
				handType: FiveOfAKind,
				hand: h,
			}
		}

		if count[i] == 4 {
			return Set{
				handType: FourOfAKind,
				hand: h,
			}
		}
	}

	// check FullHouse
	fullHouseThreeCheck := false
	fullHouseTwoCheck := false
	pairCount := 0
	for i := 0; i < numOfCards; i++ {
		if count[i] == 3 {
			fullHouseThreeCheck = true
		}

		if count[i] == 2 {
			fullHouseTwoCheck = true
			pairCount++
		}
	}

	if fullHouseThreeCheck && fullHouseTwoCheck {
		return Set{
			handType: FullHouse,
			hand: h,
		}
	}

	if fullHouseThreeCheck && !fullHouseTwoCheck {
		return Set{
			handType: ThreeOfAKind,
			hand: h,
		}
	}

	if pairCount == 2 {
		return Set{
			handType: TwoPair,
			hand: h,
		}
	}
	if pairCount == 1 {
		return Set{
			handType: OnePair,
			hand: h,
		}
	}

	return Set{
		handType: HighCard,
		hand: h,
	}
}
