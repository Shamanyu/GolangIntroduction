package main

import (
  "fmt"
)

type Player struct {
  name string
  rating int
}

func NewPlayer(name string, rating int) *Player {
  return &Player{
    name: name,
    rating: rating,
  }
}

func NewPlayerValue(name string, rating int) Player {
  return Player{
    name: name,
    rating: rating,
  }
}

func main() {
  player1 := NewPlayer("Ravi", 87)
  fmt.Println(player1)
  fmt.Println(*player1)
  player2 := NewPlayerValue("Ganguly", 86)
  fmt.Println(player2)
}
