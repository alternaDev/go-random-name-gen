package main

import (
  "testing"
  "fmt"
)

func TestGenerator(t *testing.T) {
  nameString, err := GenerateName(1, 1, 3)
  if err != nil {
    t.Error(err)
  }

  fmt.Println(nameString)
}
