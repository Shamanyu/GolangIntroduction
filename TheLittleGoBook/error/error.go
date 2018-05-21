package main

import (
  "fmt"
  "os"
  "strconv"
  "errors"
  "io"
)

func subsidiaryTester(count int) error {
  if count < 1 {
    return errors.New("Invalid count")
  }
  return nil
}

func anotherSubsidiaryTester() {
  var input int
  _, err := fmt.Scan(&input)
  if err == io.EOF {
    fmt.Println("No more input!")
  }
}

func main() {
  if len(os.Args) != 2 {
    os.Exit(1)
  }
  n, err := strconv.Atoi(os.Args[1])
  if err != nil {
    fmt.Println("Not a valid number")
  } else {
    fmt.Println(n)
  }
  subsidiaryTester(0)
  anotherSubsidiaryTester()
}
