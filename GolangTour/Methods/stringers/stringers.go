package main

import "fmt"

type Person struct {
  Name string
  Age int
}

func (p Person) String() string {
  return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
  a := Person{"Shubham Shamanyu", 24}
  z := Person{"Thor Bhai", 9001}
  fmt.Println(a, z)
}
