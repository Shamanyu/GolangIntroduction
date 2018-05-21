package main

import (
	"fmt"
)

type Pokemon struct {
	name string
	hp   int
}

func (p *Pokemon) makeEX() {
	p.hp *= 2
}

func main() {
	pikachu := Pokemon{
		name: "Pikachu",
		hp:   97,
	}
	pikachu.makeEX()
	fmt.Println(pikachu)
}
