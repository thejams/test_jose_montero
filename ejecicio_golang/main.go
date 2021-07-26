package main

import (
	"fmt"
	"poker_verification/util"
)

func main() {
	cards := []int{2, 3, 5, 7, 11, 13}
	result := util.IsStraight(cards)
	fmt.Println(result)
}
