package codility

func Solve_l7_p2(A []int) int {
	if len(A) < 2 {
		return 0
	}
	change := make([]int, len(A))
	change[0] = 0
	for i := 1; i < len(A); i++ {
		change[i] = A[i] - A[i-1]
	}

	maxprofit := 0
	maxleft := 0
	for i := 0; i < len(change); i++ {
		maxleft = max(0, maxleft+change[i])
		maxprofit = max(maxprofit, maxleft)
	}

	return maxprofit
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
