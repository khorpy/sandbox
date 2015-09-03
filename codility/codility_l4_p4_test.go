package codility

import "testing"

func TestSolve_l4_p4(t *testing.T) {
	A := []int{1, 5, 2, 1, 4, 0}

	if Solve_l4_p4(A) != 11 {
		t.Error("Return value of the solution is not correct!")
	}
}
