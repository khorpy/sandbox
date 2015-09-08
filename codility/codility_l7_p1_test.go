package codility

import "testing"

func TestSolve_l7_p1(t *testing.T) {

	A := []int{3, 2, 6, -1, 4, 5, -1, 2}

	if Solve_l7_p1(A) != 17 {
		t.Error("Return value of the solution is not correct! Expected 17, got", Solve_l7_p1(A))
	}

	A = []int{5, 5, 5}

	if Solve_l7_p1(A) != 0 {
		t.Error("Return value of the solution is not correct! Expected 0, got", Solve_l7_p1(A))
	}

	A = []int{5, 17, 0, 3}

	if Solve_l7_p1(A) != 17 {
		t.Error("Return value of the solution is not correct! Expected 17, got", Solve_l7_p1(A))
	}

}
