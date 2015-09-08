package codility

import "testing"

func TestSolve_l6_p2(t *testing.T) {

	A := []int{4, 3, 4, 4, 4, 2}

	if Solve_l6_p2(A) != 2 {
		t.Error("Return value of the solution is not correct!")
	}
}
