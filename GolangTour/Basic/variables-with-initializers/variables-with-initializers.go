package main

import "fmt"

var i, j int = 1, 2
var f1, f2 float64 = 3, 4
var r1, r2 rune = 'a', 47

func main() {
	var c, python, java, ruby = true, false, "no!", 'A'
	fmt.Println(i, j, f1, f2, r1, r2, c, python, java, ruby)
}
