package main

import (
  "fmt"
  "time"
  "math/rand"
)

func main() {
  rand.Seed(time.Now().UTC().UnixNano())
  fmt.Println("My favorite number is", rand.Intn(10))
  fmt.Println("Oh sorry, I have now changed my mind to", rand.Intn(10))
}
