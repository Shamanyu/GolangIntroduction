package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	fmt.Println("My favorite number is", rand.Intn(10))
}