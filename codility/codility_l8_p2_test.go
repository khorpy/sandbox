package codility

import "testing"

func TestSolve_l8_p2(t *testing.T) {

	if Solve_l8_p2(30) != 22 {
		t.Error("Return value of the solution is not correct! Expected 22, got ", Solve_l8_p2(30))
	}

	if Solve_l8_p2(1) != 4 {
		t.Error("Return value of the solution is not correct! Expected 4, got ", Solve_l8_p2(1))
	}

	if Solve_l8_p2(2) != 6 {
		t.Error("Return value of the solution is not correct! Expected 6, got ", Solve_l8_p2(2))
	}

	if Solve_l8_p2(36) != 24 {
		t.Error("Return value of the solution is not correct! Expected 24, got ", Solve_l8_p2(36))
	}

}
