package main

import "fmt"

type Vertex struct {
  Lat, Long float64
}

var m = map[string]Vertex{
  "Bell Labs": {40.21, -74.12},
  "Google": {37.61, -122.34},
}

func main() {
  fmt.Println(m)
}
