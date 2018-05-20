package main

import "fmt"

func split(sum int) (x, y, z int, f1, f2 float64) {
  x = sum*4/9
  y = sum*2/9
  z = sum-x-y
  f1 = float64(sum)*4/9
  f2 = float64(sum)-f1
  return
}

func main() {
  fmt.Println(split(17))
}
