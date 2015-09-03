package codility

import "testing"

func TestSolve_l4_p2(t *testing.T) {
	A := []int{10, 2, 5, 1, 8, 20}

	if Solve_l4_p2(A) != 1 {
		t.Error("Return value of the solution is not correct!")
	}

	A = []int{10, 50, 5, 1}

	if Solve_l4_p2(A) != 0 {
		t.Error("Return value of the solution is not correct!")
	}

	A = []int{5, 3, 3}

	if Solve_l4_p2(A) != 1 {
		t.Error("Return value of the solution is not correct!")
	}

}
