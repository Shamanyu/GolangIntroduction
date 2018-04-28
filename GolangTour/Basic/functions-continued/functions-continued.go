package main

import "fmt"

func add(x, y int, z float64) int {
	return x + y + int(z)
}

func main() {
	fmt.Println(add(31, 23, 21.1))
}
