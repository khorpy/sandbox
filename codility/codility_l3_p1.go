package codility

func Solve_l3_p1(A int, B int, K int) int {
	if A == 0 {
		return B/K - A/K + 1
	} else {
		return B/K - (A-1)/K
	}
}
