package main

import (
  "fmt"
  "math"
)

type I interface {
  M()
  N()
}

type T struct {
  S string
  S1 string
}

func (t *T) M() {
  fmt.Println(t.S)
}

func (t *T) N() {
  fmt.Println(t.S1)
}

type F float64

func (f F) M() {
  fmt.Println(f)
}

func (f F) N() {
  fmt.Println(f)
}

func main() {
  var i I

  i = &T{"Hello", "Bhai"}
  describe(i)
  i.M()
  i.N()

  i = F(math.Pi)
  describe(i)
  i.M()
  i.N()
}

func describe(i I) {
  fmt.Printf("(%v %T)\n", i, i)
}
