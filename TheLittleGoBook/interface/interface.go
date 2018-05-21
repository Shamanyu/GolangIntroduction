package main

import (
  // "fmt"
)

type Logger interface {
  Log(message string)
}

type Server struct {
  logger Logger
}

func process(logger Logger) {
  logger.Log("Hello!")
}

type ConsoleLogger struct {}

func main() {

}
