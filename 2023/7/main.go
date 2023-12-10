package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/jerjako/adventofcode/utils"
)

var strengthCardOrderPart1 = map[string]int{
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

type handDataPart1 struct {
	hand          string
	bid           int
	strength      int
	orderStrength int
}

func part1(lines []string) string {

	handsBid := map[string]int{}
	for _, line := range lines {
		items := strings.Split(line, " ")
		hand := strings.TrimSpace(items[0])
		bid, _ := strconv.Atoi(strings.TrimSpace(items[1]))
		handsBid[hand] = bid
	}

	hands := []handDataPart1{}
	for hand, bid := range handsBid {
		hands = append(hands, handDataPart1{
			hand:          hand,
			bid:           bid,
			strength:      calculateStrengthPart1(hand),
			orderStrength: calculateOrderStrengthPart1(hand),
		})
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].strength != hands[j].strength {
			return hands[i].strength < hands[j].strength
		}
		if hands[i].orderStrength != hands[j].orderStrength {
			return hands[i].orderStrength < hands[j].orderStrength
		}
		return strengthCardOrderPart1[hands[i].hand[0:1]] < strengthCardOrderPart1[hands[j].hand[0:1]]
	})

	total := 0
	for i, hand := range hands {
		total += hand.bid * (i + 1)
	}
	return utils.ToString(total)
}

func duplicateStrPart1(hand string) map[string]int {
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

func isFiveOfAKingPart1(hand string) bool {
	return len(duplicateStrPart1(hand)) == 1
}

func isFourOfAKingPart1(hand string) bool {
	duplicates := duplicateStrPart1(hand)
	if len(duplicates) != 2 {
		return false
	}

	for _, cnt := range duplicates {
		return cnt == 1 || cnt == 4
	}

	return false
}

func isFullHousePart1(hand string) bool {
	duplicates := duplicateStrPart1(hand)
	if len(duplicates) != 2 {
		return false
	}

	for _, cnt := range duplicates {
		return cnt == 2 || cnt == 3
	}

	return false
}

func isThreeOfAKingPart1(hand string) bool {
	duplicates := duplicateStrPart1(hand)
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

func isTwoPairPart1(hand string) bool {
	duplicates := duplicateStrPart1(hand)
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

func isOnePairPart1(hand string) bool {
	duplicates := duplicateStrPart1(hand)
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

func calculateStrengthPart1(hand string) int {
	if isFiveOfAKingPart1(hand) {
		return 6
	}
	if isFourOfAKingPart1(hand) {
		return 5
	}
	if isFullHousePart1(hand) {
		return 4
	}
	if isThreeOfAKingPart1(hand) {
		return 3
	}
	if isTwoPairPart1(hand) {
		return 2
	}
	if isOnePairPart1(hand) {
		return 1
	}
	return 0
}

func calculateOrderStrengthPart1(hand string) int {
	return strengthCardOrderPart1[hand[0:1]]*100000000 +
		strengthCardOrderPart1[hand[1:2]]*1000000 +
		strengthCardOrderPart1[hand[2:3]]*10000 +
		strengthCardOrderPart1[hand[3:4]]*100 +
		strengthCardOrderPart1[hand[4:5]]
}

var strengthCardOrderPart2 = map[string]int{
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

type handDataPart2 struct {
	hand          string
	bid           int
	strength      int
	orderStrength int
	useJoker      bool
}

func part2(lines []string) string {
	handsBid := map[string]int{}
	for _, line := range lines {
		items := strings.Split(line, " ")
		hand := strings.TrimSpace(items[0])
		bid, _ := strconv.Atoi(strings.TrimSpace(items[1]))
		handsBid[hand] = bid
	}

	hands := []handDataPart2{}
	for hand, bid := range handsBid {
		hands = append(hands, handDataPart2{
			hand:          hand,
			bid:           bid,
			strength:      calculateStrengthPart2(hand),
			orderStrength: calculateOrderStrengthPart2(hand),
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
		return strengthCardOrderPart2[hands[i].hand[0:1]] < strengthCardOrderPart2[hands[j].hand[0:1]]
	})

	total := 0
	for i, hand := range hands {
		total += hand.bid * (i + 1)
	}

	return utils.ToString(total)
}

func duplicateStrPart2(hand string) (map[string]int, bool) {
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

func isFiveOfAKingPart2(hand string) bool {
	duplicates, joker := duplicateStrPart2(hand)
	return len(duplicates) == 1 || (len(duplicates) == 2 && joker)
}

func isFourOfAKingPart2(hand string) bool {
	duplicates, joker := duplicateStrPart2(hand)
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

func isFullHousePart2(hand string) bool {
	duplicates, joker := duplicateStrPart2(hand)
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

func isThreeOfAKingPart2(hand string) bool {
	duplicates, _ := duplicateStrPart2(hand)
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

func isTwoPairPart2(hand string) bool {
	duplicates, joker := duplicateStrPart2(hand)
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

func isOnePairPart2(hand string) bool {
	duplicates, joker := duplicateStrPart2(hand)
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

func calculateStrengthPart2(hand string) int {
	if isFiveOfAKingPart2(hand) {
		return 6
	}
	if isFourOfAKingPart2(hand) {
		return 5
	}
	if isFullHousePart2(hand) {
		return 4
	}
	if isThreeOfAKingPart2(hand) {
		return 3
	}
	if isTwoPairPart2(hand) {
		return 2
	}
	if isOnePairPart2(hand) {
		return 1
	}
	return 0
}

func calculateOrderStrengthPart2(hand string) int {
	return strengthCardOrderPart2[hand[0:1]]*100000000 +
		strengthCardOrderPart2[hand[1:2]]*1000000 +
		strengthCardOrderPart2[hand[2:3]]*10000 +
		strengthCardOrderPart2[hand[3:4]]*100 +
		strengthCardOrderPart2[hand[4:5]]
}

func main() {
	lines, doPart1, doPart2 := utils.RunDay()
	if doPart1 {
		fmt.Println("result: ", part1(lines))
	}
	if doPart2 {
		fmt.Println("result: ", part2(lines))
	}
}
