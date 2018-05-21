package main

import (
	"fmt"
)

func main() {
	var scores [10]int
	scores[0] = 339
	moreScores := [4]int{9, 1, 4, 10000000}
	fmt.Println(scores, moreScores)
	for _, value := range scores {
		value *= 2
	}
	fmt.Println(scores)
	for index, _ := range scores {
		scores[index] *= 2
	}
	fmt.Println(scores)
}
