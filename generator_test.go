package nameGen

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

func BenchmarkGenerator(b *testing.B) {
  for n := 0; n < b.N; n++ {
    nameString, err := GenerateName(1, 1, 3)
    if err != nil {
      b.Error(err)
    }

    fmt.Println(nameString)
  }
}

func TestPossibilities(t *testing.T) {
  fmt.Println(GetPossibilities(1, 1, 3))
}
