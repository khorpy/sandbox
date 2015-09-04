package codility

import "testing"

func TestSolve_l5_p4(t *testing.T) {
	A := []int{4, 3, 2, 1, 5}
	B := []int{0, 1, 0, 0, 0}

	if Solve_l5_p4(A, B) != 2 {
		t.Error("Return value of the solution is not correct!", Solve_l5_p4(A, B))
	}
}
