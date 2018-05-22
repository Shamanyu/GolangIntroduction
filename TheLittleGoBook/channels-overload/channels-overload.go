package main

import (
  "fmt"
  "time"
  "math/rand"
)

type Worker struct {
  id int
}

func (w *Worker) process(c chan int) {
  for {
    data := <-c
    fmt.Printf("Worker %d got %d\n", w.id, data)
    time.Sleep(time.Millisecond*1000)
  }
}

func main() {
  c := make(chan int)
  for i:=0; i<4; i++ {
    worker := &Worker{id: i}
    go worker.process(c)
  }
  for {
    c <- rand.Int()
    fmt.Printf("Length of channel is %d\n", len(c))
    time.Sleep(time.Millisecond*10)
  }
}
