package main

import (
  "fmt"
)

func main() {
  stra := "The spice must flow"
  byts := []byte(stra)
  fmt.Println(stra, byts, string(byts))
}
