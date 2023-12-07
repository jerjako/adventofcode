package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var strengthCardOrder = map[string]int{
	"J": 0,
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"Q": 10,
	"K": 11,
	"A": 12,
}

type handData struct {
	hand          string
	bid           int
	strength      int
	orderStrength int
	useJoker      bool
}

func main() {
	chars, err := os.ReadFile("./2023/7-2/input.txt")
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
			useJoker:      strings.Contains(hand, "J"),
		})
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].strength != hands[j].strength {
			return hands[i].strength < hands[j].strength
		}
		//fmt.Println(hands[i].hand, hands[j].hand)
		/*if hands[i].useJoker && !hands[j].useJoker {
			return false
		}
		if !hands[i].useJoker && hands[j].useJoker {
			return true
		}*/
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

func duplicateStr(hand string) (map[string]int, bool) {
	duplicates := map[string]int{}
	for _, s := range hand {
		if _, ok := duplicates[string(s)]; ok {
			duplicates[string(s)]++
		} else {
			duplicates[string(s)] = 1
		}
	}

	return duplicates, strings.Contains(hand, "J")
}

func isFiveOfAKing(hand string) bool {
	duplicates, joker := duplicateStr(hand)
	return len(duplicates) == 1 || (len(duplicates) == 2 && joker)
}

func isFourOfAKing(hand string) bool {
	duplicates, joker := duplicateStr(hand)
	if len(duplicates) != 2 && (joker && len(duplicates) != 3) {
		return false
	}

	maxLen := 0
	lenJoker := 0
	for l, cnt := range duplicates {
		if l == "J" {
			lenJoker = cnt
		} else if cnt > maxLen {
			maxLen = cnt
		}
	}

	return maxLen+lenJoker == 4
}

func isFullHouse(hand string) bool {
	duplicates, joker := duplicateStr(hand)
	if len(duplicates) != 2 && (joker && len(duplicates) != 3) {
		return false
	}

	minLen := 1000
	maxLen := 0
	lenJoker := 0
	for l, cnt := range duplicates {
		if l == "J" {
			lenJoker = cnt
			continue
		} else if cnt > maxLen {
			maxLen = cnt
		}
		if cnt < minLen {
			minLen = cnt
		}
	}

	return maxLen+lenJoker == 3 && minLen == 2
}

func isThreeOfAKing(hand string) bool {
	duplicates, _ := duplicateStr(hand)
	if len(duplicates) != 3 && len(duplicates) != 4 {
		return false
	}

	minLen := 1000
	maxLen := 0
	lenJoker := 0
	for l, cnt := range duplicates {
		if l == "J" {
			lenJoker = cnt
			continue
		} else if cnt > maxLen {
			maxLen = cnt
		}
		if cnt < minLen {
			minLen = cnt
		}
	}

	return maxLen+lenJoker == 3 && minLen == 1
}

func isTwoPair(hand string) bool {
	duplicates, joker := duplicateStr(hand)
	if len(duplicates) != 3 && (joker && len(duplicates) != 4) {
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

	if pairs == 1 && joker && len(duplicates) == 4 {
		return true
	}

	return false
}

func isOnePair(hand string) bool {
	duplicates, joker := duplicateStr(hand)
	if len(duplicates) != 4 && (joker && len(duplicates) != 5) {
		return false
	}

	if joker && len(duplicates) == 5 {
		return true
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
