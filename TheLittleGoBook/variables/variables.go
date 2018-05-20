package main

import (
  "fmt"
)

func getBreakee() string {
  return "Sankar behen"
}

func main() {
  var power int
  power = 9134
  meanPower := 9134.43
  var variancePower float64
  fmt.Printf("Annoyance is over %d, with a mean of %f and variance tending to %f\n",
    power, meanPower, variancePower)
  toBreakUpWith := getBreakee()
  fmt.Printf("It's over %s\n", toBreakUpWith)
}
