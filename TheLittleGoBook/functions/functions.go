package main

import (
  "fmt"
)

func log(message string) {
  fmt.Println(message)
}

func add(a int, b int) int {
  return a+b
}

func power(name string) (int, bool) {
  if name == "Goku" {
    return 900, true
  } else {
    return -1, false
  }
}

func powerPrint(name string, power int, exists bool) {
  if exists {
    fmt.Printf("Bhai, %s ka power %d hai\n", name, power)
  } else {
    log("Bhai nahi pata inka power kitna hai")
  }
}

func main() {
  log("Bhai ye message print kardoge terminal pe?")
  sum := add(2, 3)
  fmt.Println("Bhai sum hai", sum)
  name := "Goku"
  value, exists := power(name)
  powerPrint(name, value, exists)
  name = "Shamanyu"
  value, exists = power(name)
  powerPrint(name, value, exists)
}
