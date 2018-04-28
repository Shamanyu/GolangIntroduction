package main

import "fmt"

func split(sum int) (x, y, z int) {
  x = sum*4/9
  y = sum*2/9
  z = sum-x-y
  return
}

func main() {
  fmt.Println(split(17))
}
