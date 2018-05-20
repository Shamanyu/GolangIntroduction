package main

import (
  "fmt"
)

func main() {
  name, power := "Shamanyu", 9000
  fmt.Printf("%s bhai ka power is over %d\n", name, power)
  name, power, resistance := "LameShamanyu", 9001, 56.1
  fmt.Printf("%s bhai ka power is less than %d, but resistance bhi hai %f\n",
    name, power, resistance)
}
