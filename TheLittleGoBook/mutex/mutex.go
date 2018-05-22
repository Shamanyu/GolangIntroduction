package main

import (
  "fmt"
  "time"
  "sync"
)

var (
  counter = 0
  lock sync.Mutex
)

func incr() {
  lock.Lock()
  defer lock.Unlock()
  counter++
  fmt.Println(counter)
}

func main() {
  for i:=0; i<20; i++ {
    go incr()
  }
  time.Sleep(time.Millisecond*10)
}
