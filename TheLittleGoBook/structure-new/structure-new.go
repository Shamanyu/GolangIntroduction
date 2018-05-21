package main

import (
	"fmt"
)

type lamePerson struct {
	name            string
	levelOfLameness int
}

func main() {
	doshi := new(lamePerson)
	flyingParrot := &lamePerson{}
	fmt.Println(*doshi, *flyingParrot)
}
