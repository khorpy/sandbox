package codility

import "testing"

func TestSolve_l2_p3(t *testing.T) {
	A := []int{1, 3, 6, 4, 1, 2}

	if Solve_l2_p3(A) != 5 {
		t.Error("Return value of the solution is not correct!")
	}
}
