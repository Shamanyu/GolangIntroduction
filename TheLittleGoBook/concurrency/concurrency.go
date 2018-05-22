package main

import (
  "fmt"
  "time"
)

var counter = 0

func incr() {
  counter++
  fmt.Println(counter)
}

func main() {
  for i:=0; i<20; i++ {
    go incr()
  }
  time.Sleep(time.Millisecond*10)
}
