package codility

import "testing"

func TestSolve_l7_p3(t *testing.T) {

	A := []int{3, 2, -6, 4, 0}

	if Solve_l7_p3(A) != 5 {
		t.Error("Return value of the solution is not correct! Expected 5, got", Solve_l7_p3(A))
	}

	A = []int{-2, 1}

	if Solve_l7_p3(A) != 1 {
		t.Error("Return value of the solution is not correct! Expected 1, got", Solve_l7_p3(A))
	}

	A = []int{-5, -5, -5, -5}

	if Solve_l7_p3(A) != -5 {
		t.Error("Return value of the solution is not correct! Expected -5, got", Solve_l7_p3(A))
	}

}
