package codility

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
				a[i], err = strconv.Atoi(values[i])
				if err != nil {
					panic("Error reading input: can not convert to int")
				}
			}
			fmt.Println(solve(a))
		}
	}
}

func solve(a []int) int {
	var sum int = 0
	var sum_exp int = len(a) + 1

	for i, v := range a {
		sum += v
		sum_exp += i + 1
	}
	return sum_exp - sum
}
