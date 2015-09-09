package codility

import "testing"

func TestSolve_l7_p2(t *testing.T) {

	A := []int{23171, 21011, 21123, 21366, 21013, 21367}

	if Solve_l7_p2(A) != 356 {
		t.Error("Return value of the solution is not correct! Expected 17, got", Solve_l7_p1(A))
	}
}
