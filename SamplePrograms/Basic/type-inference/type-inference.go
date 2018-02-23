package main

import "fmt"

func main() {
	v1 := 42
	v2 := 3.142
	v3 := 0.867 + 0.5i
	fmt.Printf("v1, v2, v3 are of types %T %T %T\n", v1, v2, v3)
}