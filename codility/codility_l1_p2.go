package codility

func Solve_l1_p2(x, y, d int) int {
	min := (y - x) / d
	if (y-x)%d > 0 {
		min++
	}
	return min
}
