package codility

import "testing"

func TestSolve_l2_p2(t *testing.T) {
	A := []int{4, 1, 3, 5, 2}

	if Solve_l2_p2(A) != 1 {
		t.Error("Return value of the solution is not correct!")
	}

	A = []int{4, 1, 3, 2, 6}

	if Solve_l2_p2(A) != 0 {
		t.Error("Return value of the solution is not correct!")
	}
}
