package main

import "fmt"

type I interface {
  M()
  N()
}

type T struct {
  S string
}

func (t T) M() {
  fmt.Println(t.S)
}

func (t T) N() {
  fmt.Println(t)
}

func main() {
  var i I = T{"hello"}
  i.M()
  i.N()
}
