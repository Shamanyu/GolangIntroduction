package main

import "fmt"

func main() {
  doorKeeperName := "Shamanyu"
  if doorKeeperName == "Leto" {
    fmt.Println("Avoid entering the door")
  } else {
    fmt.Println("Please say hi to", doorKeeperName)
  }
}
