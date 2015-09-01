package codility

import "testing"

func TestSolve_l3_p3(t *testing.T) {
	A := []int{4, 2, 2, 5, 1, 5, 8}

	if Solve_l3_p3(A) != 1 {
		t.Error("Return value of the solution is not correct!")
	}
}
