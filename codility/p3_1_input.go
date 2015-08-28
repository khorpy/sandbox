package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
)

const (
	cases_max int = 1000
	N_max     int = 100000
)

var cases_to_gen int
var N int
var values []string

func shuffle(a []int) {
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

func init() {
	flag.IntVar(&cases_to_gen, "c", 0, "number of test cases to generate")
	flag.IntVar(&N, "N", 0, "N value")
}

func main() {
	flag.Parse()
	if cases_to_gen <= 0 || cases_to_gen > cases_max || N <= 0 || N > N_max {
		return
	}
	values = make([]string, N)
	for i := 0; i < cases_to_gen; i++ {
		a := make([]int, 0)
		elem_miss := 1 + rand.Int()%N
		for j := 1; j <= N+1; j++ {
			if j != elem_miss {
				a = append(a, j)
			}
		}
		shuffle(a)
		for j, v := range a {
			values[j] = fmt.Sprint(v)
		}
		fmt.Println(strings.Join(values, " "))
	}
}
