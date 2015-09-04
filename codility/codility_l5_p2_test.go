package codility

import "testing"

func TestSolve_l5_p2(t *testing.T) {
	H := []int{8, 8, 5, 7, 9, 8, 7, 4, 8}

	if Solve_l5_p2(H) != 7 {
		t.Error("Return value of the solution is not correct!")
	}
}
