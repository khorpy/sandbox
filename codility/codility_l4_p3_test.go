package codility

import "testing"

func TestSolve_l4_p3(t *testing.T) {
	A := []int{10, 2, 5, 1, 2, 20}

	if Solve_l4_p3(A) != 5 {
		t.Error("Return value of the solution is not correct!")
	}
}
