package main

import (
  "fmt"
)

type Programmer struct {
  name string
  proficiency int
  mentor *Programmer
}

func main() {
  shubham := &Programmer{
    name: "Shubham",
    proficiency: 3,
    mentor: &Programmer{
      name: "Doshi",
      proficiency: 5,
      mentor: nil,
    },
  }
  fmt.Println(shubham)
}
