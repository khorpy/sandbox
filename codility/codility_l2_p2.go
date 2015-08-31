package codility

func Solve_l2_p2(A []int) int {
	var total int = 0
	counters := make([]int, len(A)+1)
	for i := 0; i < len(counters); i++ {
		counters[i] = 0
	}
	for i := 0; i < len(A); i++ {
		if A[i] < len(counters) {
			if counters[A[i]] == 0 {
				counters[A[i]]++
				total++
			}
		}
	}
	if total == len(A) {
		return 1
	} else {
		return 0
	}
}
