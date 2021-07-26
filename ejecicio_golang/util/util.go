package util

import (
	"sort"
)

// Cherk if all cards are valid
func AreAllValidCards(s []int) bool {
	for _, card := range s {
		if card < 1 || card > 14 {
			return false
		}
	}

	return true
}

// SortAndVerifyCards Sorts the slice of int and check for the number 14 transformation
func SortAndVerifyCards(s []int) []int {
	for _, val := range s {
		if val == 14 {
			s = append(s, 1)
		}
	}
	sort.Ints(s)

	return s
}

// IsStraight Cherk if a hand has a straight flush
func IsStraight(cards []int) bool {
	if len(cards) > 7 || len(cards) < 5 {
		return false
	}

	cards = SortAndVerifyCards(cards)

	if !AreAllValidCards(cards) {
		return false
	}

	lastIndex := cards[0]
	straightCount := 1
	for index, value := range cards {
		if index == 0 {
			continue
		}
		if lastIndex+1 == value {
			straightCount += 1
		} else {
			straightCount = 1
		}

		lastIndex = value

		if straightCount == 5 {
			return true
		}
	}

	return false
}
