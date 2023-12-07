package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var strengthCardOrder = map[string]int{
	"2": 0,
	"3": 1,
	"4": 2,
	"5": 3,
	"6": 4,
	"7": 5,
	"8": 6,
	"9": 7,
	"T": 8,
	"J": 9,
	"Q": 10,
	"K": 11,
	"A": 12,
}

type handData struct {
	hand          string
	bid           int
	strength      int
	orderStrength int
}

func main() {
	chars, err := os.ReadFile("./2023/7/input.txt")
	if err != nil {
		panic(err)
	}

	handsBid := map[string]int{}
	for _, line := range strings.Split(string(chars), "\n") {
		items := strings.Split(line, " ")
		hand := strings.TrimSpace(items[0])
		bid, _ := strconv.Atoi(strings.TrimSpace(items[1]))
		handsBid[hand] = bid
	}

	hands := []handData{}
	for hand, bid := range handsBid {
		hands = append(hands, handData{
			hand:          hand,
			bid:           bid,
			strength:      calculateStrength(hand),
			orderStrength: calculateOrderStrength(hand),
		})
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].strength != hands[j].strength {
			return hands[i].strength < hands[j].strength
		}
		if hands[i].orderStrength != hands[j].orderStrength {
			return hands[i].orderStrength < hands[j].orderStrength
		}
		return strengthCardOrder[hands[i].hand[0:1]] < strengthCardOrder[hands[j].hand[0:1]]
	})

	total := 0
	for i, hand := range hands {
		total += hand.bid * (i + 1)
	}
	fmt.Println("total: ", total)
}

func duplicateStr(hand string) map[string]int {
	duplicates := map[string]int{}
	for _, s := range hand {
		if _, ok := duplicates[string(s)]; ok {
			duplicates[string(s)]++
		} else {
			duplicates[string(s)] = 1
		}
	}

	return duplicates
}

func isFiveOfAKing(hand string) bool {
	return len(duplicateStr(hand)) == 1
}

func isFourOfAKing(hand string) bool {
	duplicates := duplicateStr(hand)
	if len(duplicates) != 2 {
		return false
	}

	for _, cnt := range duplicates {
		return cnt == 1 || cnt == 4
	}

	return false
}

func isFullHouse(hand string) bool {
	duplicates := duplicateStr(hand)
	if len(duplicates) != 2 {
		return false
	}

	for _, cnt := range duplicates {
		return cnt == 2 || cnt == 3
	}

	return false
}

func isThreeOfAKing(hand string) bool {
	duplicates := duplicateStr(hand)
	if len(duplicates) != 3 {
		return false
	}

	for _, cnt := range duplicates {
		if cnt == 3 {
			return true
		}
	}

	return false
}

func isTwoPair(hand string) bool {
	duplicates := duplicateStr(hand)
	if len(duplicates) != 3 {
		return false
	}

	pairs := 0
	for _, cnt := range duplicates {
		if cnt == 2 {
			pairs++
			if pairs == 2 {
				return true
			}
		}
	}

	return false
}

func isOnePair(hand string) bool {
	duplicates := duplicateStr(hand)
	if len(duplicates) != 4 {
		return false
	}

	pairs := 0
	for _, cnt := range duplicates {
		if cnt == 2 {
			pairs++
		}
	}

	return pairs == 1
}

func calculateStrength(hand string) int {
	if isFiveOfAKing(hand) {
		return 6
	}
	if isFourOfAKing(hand) {
		return 5
	}
	if isFullHouse(hand) {
		return 4
	}
	if isThreeOfAKing(hand) {
		return 3
	}
	if isTwoPair(hand) {
		return 2
	}
	if isOnePair(hand) {
		return 1
	}
	return 0
}

func calculateOrderStrength(hand string) int {
	return strengthCardOrder[hand[0:1]]*100000000 +
		strengthCardOrder[hand[1:2]]*1000000 +
		strengthCardOrder[hand[2:3]]*10000 +
		strengthCardOrder[hand[3:4]]*100 +
		strengthCardOrder[hand[4:5]]
}
