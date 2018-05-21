package main

import (
  "fmt"
)

type Pokemon struct {
  name string
  hp int
}

func (p *Pokemon) Introduce() {
  fmt.Printf("Hi, I am %s\n", p.name)
}

type AncientPokemon struct {
  *Pokemon
  age int
}

func main() {
  articuno := &AncientPokemon{
    Pokemon: &Pokemon{
      name: "Articuno",
      hp: 97,
    },
    age: 20001,
  }
  articuno.Introduce()
  fmt.Println(*articuno)
  fmt.Println(articuno.name, articuno.Pokemon.name)
}
