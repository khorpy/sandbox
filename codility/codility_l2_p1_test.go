package codility

import "testing"

func TestSolve_l2_p1(t *testing.T) {
	X := 5
	A := []int{1, 3, 1, 4, 2, 3, 5, 4}

	if Solve_l2_p1(X, A) != 6 {
		t.Error("Return value of a solution is not correct!")
	}

	X = 5
	A = []int{1, 3, 1, 4, 5, 3, 5, 4}

	if Solve_l2_p1(X, A) != -1 {
		t.Error("Return value of a solution is not correct!")
	}
}
