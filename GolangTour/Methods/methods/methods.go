package main

import (
  "fmt"
  "math"
)

type Vertex struct {
  X, Y, Z float64
}

func (v Vertex) Abs() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z)
}

func (v Vertex) Scaler(factor float64) float64 {
  return v.Abs()*factor
}

func main() {
  v := Vertex{3, 4, 10}
  fmt.Println(v.Abs())
  fmt.Println(v.Scaler(2))
}
