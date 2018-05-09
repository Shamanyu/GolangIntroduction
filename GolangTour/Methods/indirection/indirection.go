package main

import "fmt"

type Vertex struct {
  X, Y, Z float64
}

func (v *Vertex) Scale(f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
  v.Z = v.Z * f
}

func ScaleFunc(v *Vertex, f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
  v.Z = v.Z * f
}

func main() {
  v := Vertex{3, 4, 5}
  (&v).Scale(2)
  v.Scale(2)
  ScaleFunc(&v, 10)

  p := &Vertex{4, 3, 10}
  p.Scale(3)
  ScaleFunc(p, 8)

  fmt.Println(v, p)
}
