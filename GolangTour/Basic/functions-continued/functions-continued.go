package main

import "fmt"

func add(x, y int, z, w float64) int {
	return x + y + int(z) + int(w)
}

func main() {
	fmt.Println(add(31, 23, 21.1, 43.0))
}
