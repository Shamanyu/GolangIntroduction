package main

import (
  "fmt"
)

func subsidiaryTester() {
  scores := []int{1, 2, 3, 4, 5}
  slice := scores[2:4]
  slice[0] = 99
  fmt.Println(scores)
}

func removeAtIndex(source []int, index int) []int {
  lastIndex := len(source) - 1
  source[index], source[lastIndex] = source[lastIndex], source[index]
  return source[:lastIndex]
}

func subsidiaryTesterTwo() {
  scores := []int{1, 2, 3, 4, 5}
  scores = removeAtIndex(scores, 2)
  fmt.Println(scores)
}

func main() {
  scores := []int{1, 4, 9, 16}
  moreScores := make([]int, 10)
  evenMoreScores := make([]int, 0, 10)
  fmt.Println(scores, moreScores, evenMoreScores, len(scores), len(moreScores),
    len(evenMoreScores))
  ughScoresAgain := make([]int, 0, 10)
  ughScoresAgain = append(ughScoresAgain, 5)
  fmt.Println(ughScoresAgain)
  yuckScores := make([]int, 0, 10)
  yuckScores = yuckScores[0:8]
  yuckScores[7] = 43
  fmt.Println(yuckScores)
  gameTwoScores := make([]int, 5)
  gameTwoScores = append(gameTwoScores, 89)
  fmt.Println(gameTwoScores, len(gameTwoScores), cap(gameTwoScores))
  gameThreeScores := make([]int, 2, 10)
  gameThreeScores = append(gameThreeScores, 91)
  fmt.Println(gameThreeScores, len(gameThreeScores), cap(gameThreeScores))
  peopleOne := []string{"Shamanyu", "Sis", "Behen"}
  checks := make([]bool, 10)
  var peopleTwo []string
  finalScores := make([]int, 0, 20)
  fmt.Println(peopleOne, len(peopleOne), cap(peopleOne), checks, len(checks),
    cap(checks), peopleTwo, len(peopleTwo), cap(peopleTwo), finalScores,
    len(finalScores), cap(finalScores))
  subsidiaryTester()
  subsidiaryTesterTwo()
}
