package codility

import "testing"

func TestSolve_l8_p1(t *testing.T) {

	if Solve_l8_p1(24) != 8 {
		t.Error("Return value of the solution is not correct! Expected 8, got ", Solve_l8_p1(24))
	}

	if Solve_l8_p1(1) != 1 {
		t.Error("Return value of the solution is not correct! Expected 1, got ", Solve_l8_p1(1))
	}

	if Solve_l8_p1(16) != 5 {
		t.Error("Return value of the solution is not correct! Expected 5, got ", Solve_l8_p1(16))
	}

}
