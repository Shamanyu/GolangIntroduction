package main

import (
  "fmt"
  "time"
)

func main() {
  tick1 := time.Tick(1000 * time.Millisecond)
  tick2 := time.Tick(1000 * time.Millisecond)
  boom := time.Tick(5500 * time.Millisecond)
  for {
    select {
    case <-tick1:
      fmt.Println("tick1.")
    case <-tick2:
      fmt.Println("tick2.")
    case <-boom:
      fmt.Println("boom.")
    default:
      fmt.Println("  .")
      time.Sleep(50 * time.Millisecond)
    }
  }
}
