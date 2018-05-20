package main

import (
  "fmt"
)

type Pokemon struct {
  name string
  hp int
}

func makeEXClone(p Pokemon) {
  p.hp *= 2
}

func makeEXOriginal(p *Pokemon) {
  p.hp *= 2
}

func main() {
  pichi := Pokemon {
    name: "Pichi",
    hp: 76,
  }
  fmt.Println(pichi)
  makeEXClone(pichi)
  fmt.Println(pichi)
  makeEXOriginal(&pichi)
  fmt.Println(pichi)
}
