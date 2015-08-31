package codility

func Solve_l2_p4(N int, A []int) []int {
	counters := make([]int, N)
	curr_max := 0
	curr_min := 0
	for i := 0; i < len(A); i++ {
		if A[i] == N+1 {
			curr_min = curr_max
		} else {
			counters[A[i]-1] = max(counters[A[i]-1], curr_min) + 1
			if counters[A[i]-1] > curr_max {
				curr_max = counters[A[i]-1]
			}
		}
	}

	for i := 0; i < len(counters); i++ {
		counters[i] = max(counters[i], curr_min)
	}

	return counters
}

func max(x, y int) int {
	if y > x {
		return y
	} else {
		return x
	}
}
