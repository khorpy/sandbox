package codility

func Solve_l3_p2(A []int) int {
	sum, passing := 0, 0
	for _, v := range A {
		sum += v
	}
	for _, v := range A {
		if v == 0 {
			passing += sum
			if passing > 1000000000 {
				return -1
			}
		} else {
			sum--
		}
	}
	return passing
}
