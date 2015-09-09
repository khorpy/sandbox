package codility

import "testing"

func TestSolve_l8_p3(t *testing.T) {

	A := []int{1, 2, 3, 4, 3, 4, 1, 2, 3, 4, 6, 2}

	if Solve_l8_p3(A) != 3 {
		t.Error("Return value of the solution is not correct! Expected 3, got ", Solve_l8_p3(A))
	}

	A = []int{1, 1, 1, 1, 1, 1, 1}

	if Solve_l8_p3(A) != 0 {
		t.Error("Return value of the solution is not correct! Expected 0, got ", Solve_l8_p3(A))
	}

	A = []int{0, 1, 0, 0, 1, 0, 0, 1, 0}

	if Solve_l8_p3(A) != 3 {
		t.Error("Return value of the solution is not correct! Expected 3, got ", Solve_l8_p3(A))
	}

	A = []int{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1}

	if Solve_l8_p3(A) != 8 {
		t.Error("Return value of the solution is not correct! Expected 8, got ", Solve_l8_p3(A))
	}

}
