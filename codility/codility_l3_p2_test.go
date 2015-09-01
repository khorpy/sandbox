package codility

import "testing"

func TestSolve_l3_p2(t *testing.T) {
	A := []int{0, 1, 0, 1, 1}

	if Solve_l3_p2(A) != 5 {
		t.Error("Return value of the solution is not correct!")
	}
}
