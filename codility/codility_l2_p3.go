package codility

func Solve_l2_p3(A []int) int {
	counters := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		if A[i] > 0 && A[i] <= len(A) {
			counters[A[i]-1] += 1
		}
	}
	i := 0
	for i < len(counters) {
		if counters[i] == 0 {
			break
		}
		i++
	}
	return i + 1
}
