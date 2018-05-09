package main

import (
  "fmt"
  "time"
  "strconv"
)

func say(s string) {
  for i := 0; i < 5; i++{
    time.Sleep(100 * time.Millisecond)
    fmt.Println(s + " " + strconv.Itoa(i))
  }
}

func main() {
  go say("world")
  say ("hello")
  // time.Sleep(1000 * time.Millisecond)
}
