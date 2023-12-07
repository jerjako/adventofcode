package main

import (
	"testing"
)

/*
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
}*/

func Test_calculateStrength(t *testing.T) {
	handsTest := map[string]int{
		"K6TTT": 3,
		"QTTQT": 4,
		"T6974": 0,
		"KA396": 0,
		"22A22": 5,
		"73629": 0,
		"3K33J": 5,
		"767Q6": 2,
		"A88J8": 5,
		"28789": 1,
		"6JT83": 1,
		"72629": 1,
		"A7J2K": 1,
		"K4879": 0,
		"QK76J": 1,
		"TJ923": 1,
		"J6664": 5,
		"4QTQQ": 3,
		"9K476": 0,
		"J7A77": 5,
		"8KK87": 2,

		"AAA8A": 5,
		"JAJKQ": 3,
		"97J77": 5,
		"K7K9J": 3,
		"88JJA": 5,
		"74A44": 3,
		"KKAJK": 5,
		"3TT5Q": 1,
		"27762": 2,
		"TT3TT": 5,
		"6388Q": 1,
		"3T48T": 1,
		"AJQJ5": 3,
		"65K2Q": 0,
		"A4652": 0,
		"65A72": 0,
		"J66J6": 6,
		"44434": 5,
		"99777": 4,
		"3J2Q2": 3,
		"66QQ8": 2,
		"A999J": 5,
		"AAJJ2": 5,

		"2345A": 0,
		"Q2KJJ": 3,
		"Q2Q2Q": 4,
		"T3T3J": 4,
		"T3Q33": 3,
		"2345J": 1,
		"J345A": 1,
		"32T3K": 1,
		"T55J5": 5,
		"KK677": 2,
		"KTJJT": 5,
		"QQQJA": 5,
		"JJJJJ": 6,
		"JAAAA": 6,
		"AAAAJ": 6,
		"AAAAA": 6,
		"2AAAA": 5,
		"2JJJJ": 6,
		"JJJJ2": 6,
	}

	for hand, strength := range handsTest {
		t.Run(hand, func(t *testing.T) {
			if got := calculateStrength(hand); got != strength {
				t.Errorf("calculateStrength() = %v, want %v", got, strength)
			}
		})
	}
}
