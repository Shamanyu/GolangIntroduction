package main

import (
  "fmt"
)

func main() {
  count := 9
  if x := 10; count > x {
    fmt.Println("More")
  } else if count == x {
    fmt.Println("Equal")
  } else {
    fmt.Println("Less")
  }
}
