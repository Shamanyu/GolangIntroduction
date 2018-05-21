package main

import (
	"fmt"
)

type Saiyan struct {
	name  string
	power int
}

func main() {
	goku := Saiyan{
		name:  "Goku",
		power: 90,
	}
	shubham := Saiyan{name: "Shubham"}
	mysterySaiyan := Saiyan{
		power: 91,
	}
	shamanyu := Saiyan{name: "Shamanyu"}
	shamanyu.power = 94
	srishti := Saiyan{"Srishti", 96}
	fmt.Println(goku, shubham, mysterySaiyan, shamanyu, srishti)
}
