package codility

func Solve_l2_p1(X int, A []int) int {
	var path int = 0

	counters := make([]int, X+1)
	for i := 0; i < X; i++ {
		counters[i] = 0
	}

	i := 0
	for i < len(A) {
		if A[i] <= X {
			if counters[A[i]] == 0 {
				counters[A[i]] = 1
				path++
			}
		}
		if path == X {
			return i
		}
		i++
	}
	return -1
}
