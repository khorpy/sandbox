package codility

func Solve_l1_p1(a []int) int {
	total, left := 0, 0

	for _, v := range a {
		total += v
	}

	min := abs(2*a[0] - total)

	for i := 1; i < len(a); i++ {
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
