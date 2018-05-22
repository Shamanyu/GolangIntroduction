package main

import (
  "fmt"
  "time"
)

func process() {
  fmt.Println("Processing")
}

func main() {
  fmt.Println("Start")
  go process()
  time.Sleep(time.Millisecond * 10)
  go process()
  fmt.Println("End")
}
