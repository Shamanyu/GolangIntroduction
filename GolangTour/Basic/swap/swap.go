package main

import "fmt"

func swap(x, y string) (string, string, int) {
	return y, x, 1
}

func main() {
	var a, b string
	a, b, _ = swap("Hello", "World")
	a, c, _ := swap("Hello", "World")
	fmt.Println(a, b, c)
}
