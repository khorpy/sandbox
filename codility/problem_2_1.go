package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Not enought arguments to solve!")
		return
	}
	x, _ := strconv.Atoi(os.Args[1])
	y, _ := strconv.Atoi(os.Args[2])
	d, _ := strconv.Atoi(os.Args[3])
	fmt.Println(solve(x, y, d))
}

func solve(x, y, d int) int {
	min := (y - x) / d
	if (y-x)%d > 0 {
		min++
	}
	return min
}
