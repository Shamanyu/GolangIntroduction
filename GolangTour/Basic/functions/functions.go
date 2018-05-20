package main

import "fmt"

func add(x float64, y float64) float64 {
	return x + y
}

func subtract(x float64, y float64) float64 {
	return x - y
}

func multiply(x float64, y float64) float64 {
	return x * y
}

func divide(x float64, y float64) float64 {
	return x / y
}

func basicOperations(operator string, x float64, y float64) float64 {
	if operator == "+" {
		return add(x, y)
	} else if operator == "-" {
		return subtract(x, y)
	} else if operator == "*" {
		return multiply(x, y)
	} else if operator == "/" {
		return divide(x, y)
	} else {
		return 0
	}
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(add(23, 11))
	fmt.Println(basicOperations("+", 100, -5))
	fmt.Println(basicOperations("-", 100, -5))
	fmt.Println(basicOperations("*", 100, -5))
	fmt.Println(basicOperations("/", 100, -5))
	fmt.Println(basicOperations("/", 100, 0))
	fmt.Println(basicOperations("/", -100, 0))
	fmt.Println(basicOperations("$", 300, 0))
}
