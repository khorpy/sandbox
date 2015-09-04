package codility

func Solve_l5_p4(A []int, B []int) int {
	dq := make([]int, 0)
	count := 0
	for len(A) > 0 {
		if B[0] == 0 {
			if len(dq) > 0 {
				for len(dq) > 0 {
					if dq[len(dq)-1] > A[0] {
						A = A[1:]
						B = B[1:]
						break
					} else {
						dq = dq[:len(dq)-1]
					}
				}
			} else {
				A = A[1:]
				B = B[1:]
				count++
			}
		} else {
			dq = append(dq, A[0])
			A = A[1:]
			B = B[1:]
		}
	}
	return count + len(dq)
}
