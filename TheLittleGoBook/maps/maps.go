package main

import (
	"fmt"
)

type Saiyan struct {
	Name    string
	Friends map[string]*Saiyan
}

func structureTester() {
	goku := &Saiyan{
		Name:    "Goku",
		Friends: make(map[string]*Saiyan),
	}
	fmt.Println(goku)
}

func main() {
	lookup := make(map[string]int)
	lookup["goku"] = 9001
	power, exists := lookup["vegeta"]
	fmt.Println(power, exists)
	power, exists = lookup["goku"]
	fmt.Println(power, exists)
	fmt.Println(len(lookup))
	delete(lookup, "goku")
	fmt.Println(len(lookup))
	lookupAgain := make(map[string]int, 100)
	fmt.Println(len(lookupAgain))
	structureTester()
	finalLookup := map[string]int{
		"goku":  9001,
		"gohan": 2000,
	}
	fmt.Println(finalLookup)
}
