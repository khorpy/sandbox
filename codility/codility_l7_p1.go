package codility

func Solve_l7_p1(A []int) int {
	max_left := make([]int, len(A))
	max_end := 0
	for i := 1; i < len(A)-1; i++ {
		max_end = max(0, max_end+A[i])
		max_left[i] = max_end
	}
	max_right := make([]int, len(A))
	max_end = 0
	for i := len(A) - 2; i > 0; i-- {
		max_end = max(0, max_end+A[i])
		max_right[i] = max_end
	}
	max_double := 0
	for i := 1; i < len(A)-1; i++ {
		if max_left[i-1]+max_right[i+1] > max_double {
			max_double = max_left[i-1] + max_right[i+1]
		}
	}
	return max_double
}

/*
func max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}
*/
