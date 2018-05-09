package main

import (
  "fmt"
  "math"
)

type Vertex struct {
  X, Y, Z float64
}

func Abs(v Vertex) float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z+v.Z)
}

func main() {
  v := Vertex{3, 4, 5}
  fmt.Println(Abs(v))
}
