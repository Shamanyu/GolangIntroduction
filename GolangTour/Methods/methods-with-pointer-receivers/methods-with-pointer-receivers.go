package main

import (
  "fmt"
  "math"
)

type Vertex struct {
  X, Y, Z float64
}

func (v *Vertex) Scale(f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
  v.Z = v.Z * f
}

func (v *Vertex) Abs() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func main() {
  v := &Vertex{3, 4, 5}
  fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
  v.Scale(5)
  fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
