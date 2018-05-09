package main

import (
  "fmt"
  "math"
)

type Vertex struct {
  X, Y, Z float64
}

func (v Vertex) Abs() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v *Vertex) Scale(f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
  v.Z = v.Z * f
}

func main() {
  v := Vertex{3, 4, 20}
  v.Scale(10)
  fmt.Println(v.Abs())
}
