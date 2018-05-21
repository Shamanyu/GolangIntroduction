package main

import (
  "fmt"
)

func add(a interface{}, b interface{}) interface{} {
  switch a.(type) {
    case int:
      return a.(int)+b.(int)
    case string:
      return a.(string)+b.(string)
    default:
      return "Unsupported type"
  }
}

func main() {
  fmt.Printf("%d %s %s\n", add(2, 3), add("Shubham", "Shamanyu"), add(true, false))
}
