package main 

import (
  "fmt"
  "math/rand"
  "flag"
  "strings"
)

const (
  cases_max int = 10000
  case_size int = 100
  min_value int = -1000
  max_value int = 1000
)

var cases_to_gen int
var values []string

func init() {
  flag.IntVar(&cases_to_gen, "n", 0, "number of test cases to generate")
}

func main() {
  flag.Parse()
  values = make([]string, case_size)
  if cases_to_gen > 0 && cases_to_gen <= cases_max {
    for i := 0; i < cases_to_gen; i++ {
      for j := 0; j < case_size; j++ {
        values[j] = fmt.Sprint(1000 - rand.Int() % 2000)
      }
      fmt.Println(strings.Join(values, " "))
    }
  }
}