package codility

import "testing"

func TestSolve_l5_p3(t *testing.T) {
	S := "{[()()]}"

	if Solve_l5_p3(S) != 1 {
		t.Error("Return value of the solution is not correct!")
	}

	S = "([)()]"

	if Solve_l5_p3(S) != 0 {
		t.Error("Return value of the solution is not correct!")
	}

	S = ")("

	if Solve_l5_p3(S) != 0 {
		t.Error("Return value of the solution is not correct!")
	}
}
