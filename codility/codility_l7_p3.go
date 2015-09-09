package codility

func Solve_l7_p3(A []int) int {
	if len(A) == 1 {
		return A[0]
	}
	maxleft := A[0]
	maxsum := A[0]
	for i := 1; i < len(A); i++ {
		maxleft = max(A[i], maxleft+A[i])
		maxsum = max(maxsum, maxleft)
	}
	return maxsum
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
