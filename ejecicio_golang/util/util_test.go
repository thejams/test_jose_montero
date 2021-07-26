package util_test

import (
	"poker_verification/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsStraight(t *testing.T) {
	t.Run("should return true when a slice of cards has a straight in it", func(t *testing.T) {
		type TestCases struct {
			data     []int
			response bool
		}

		testCases := []TestCases{
			{[]int{2, 3, 5, 4, 6}, true},
			{[]int{14, 5, 4, 2, 3}, true},
			{[]int{7, 7, 12, 11, 3, 4, 14}, false},
			{[]int{7, 3, 2}, false},
			{[]int{14, 2, 3, 4, 5}, true},
			{[]int{9, 10, 11, 12, 13}, true},
			{[]int{2, 7, 8, 5, 10, 9, 11}, true},
			{[]int{7, 8, 12, 13, 14}, false},
		}

		for _, test := range testCases {
			result := util.IsStraight(test.data)
			assert.Equal(t, result, test.response)
		}
	})
}

func TestAreAllValidCards(t *testing.T) {
	t.Run("should return false when numbers not between 2 and 14 are received", func(t *testing.T) {
		cards := []int{2, 21, 5, 4, 6}
		result := util.AreAllValidCards(cards)

		assert.Equal(t, result, false)
	})

	t.Run("should return true when all numbers are between 2 and 14", func(t *testing.T) {
		cards := []int{2, 3, 5, 4, 6}
		result := util.AreAllValidCards(cards)

		assert.Equal(t, result, true)
	})
}

func BenchmarkIsStraight(b *testing.B) {
	cards := []int{2, 7, 8, 5, 10, 9, 11}
	for i := 0; i < b.N; i++ {
		util.IsStraight(cards)
	}
}

func BenchmarkAreAllValidCards(b *testing.B) {
	cards := []int{2, 7, 8, 5, 10, 9, 11}
	for i := 0; i < b.N; i++ {
		util.AreAllValidCards(cards)
	}
}
