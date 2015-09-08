package codility

import "testing"

func TestSolve_l6_p1(t *testing.T) {

	A := []int{3, 4, 3, 2, 3, -1, 3, 3}

	if Solve_l6_p1(A) != 0 && Solve_l6_p1(A) != 2 && Solve_l6_p1(A) != 4 && Solve_l6_p1(A) != 6 && Solve_l6_p1(A) != 7 {
		t.Error("Return value of the solution is not correct!")
	}
}
