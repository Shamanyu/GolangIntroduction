package main

import (
  "fmt"
  "time"
  "math"
  "math/rand"
)

func main() {
  rand.Seed(time.Now().UTC().UnixNano())
  fmt.Printf("Would you rather have %d problems or %d problems?\n", int(math.Sqrt(25)), rand.Intn(10))
}
