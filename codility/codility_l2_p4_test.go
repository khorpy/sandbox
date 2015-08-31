package codility

import "testing"

func TestSolve_l2_p4(t *testing.T) {
	N := 5
	A := []int{3, 4, 4, 6, 1, 4, 4}
	R := []int{3, 2, 2, 4, 2}

	result := Solve_l2_p4(N, A)
	if len(R) != len(result) {
		t.Error("Expected result array is of different size")
	} else {
		for i := 0; i < len(R); i++ {
			if R[i] != result[i] {
				t.Error("Expected and actual arrays do not match")
				return
			}
		}
	}
}
