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

func (ap *AncientPokemon) Introduce() {
  fmt.Printf("Hi, I am %s with an age of %d\n", ap.name, ap.age)
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
  articuno.Pokemon.Introduce()
  fmt.Println(*articuno)
  fmt.Println(articuno.name, articuno.Pokemon.name)
}
