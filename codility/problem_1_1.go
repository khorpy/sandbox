package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "strconv"
)

var s string
var a []int

func main() {
  input := bufio.NewReader(os.Stdin)
  for {
    s, err := input.ReadString('\n')
    if err != nil {
      break
    } else {
      values := strings.Split(strings.Trim(s, "\n "), " ")
      a = make([]int, len(values))
      for i := 0; i < len(values); i++ {
        v, err := strconv.ParseInt(values[i], 10, 0)
        if err != nil {
          panic("Error reading input: can not convert to int")
        }
        a[i] = int(v)
      }
      fmt.Println(solve(a))
    }
  }  
}

func solve(a []int) int {
  total, left := 0, 0

  for _, v := range a {
    total += v
  }

  min := abs(2*a[0] - total)

  for i := 1; i < len(a); i++  {
    left += a[i-1]
    diff := abs(2*left - total) 
    if diff < min {
      min = diff
    }
  }

  return min
}

func abs(x int) int {
  if x >= 0 {
    return x
  } else {
    return -x
  }
}