package codility

import "testing"

func TestSolve_l3_p1(t *testing.T) {
	A, B, K := 6, 11, 2

	if Solve_l3_p1(A, B, K) != 3 {
		t.Error("Return value of the solution is not correct!")
	}

	A, B, K = 0, 0, 5

	if Solve_l3_p1(A, B, K) != 1 {
		t.Error("Return value of the solution is not correct!")
	}

	A, B, K = 11, 14, 2

	if Solve_l3_p1(A, B, K) != 2 {
		t.Error("Return value of the solution is not correct!")
	}

	A, B, K = 0, 2000000000, 1

	if Solve_l3_p1(A, B, K) != 2000000001 {
		t.Error("Return value of the solution is not correct!")
	}

	A, B, K = 11, 345, 17

	if Solve_l3_p1(A, B, K) != 20 {
		n := 0
		for i := A; i <= B; i++ {
			if i%K == 0 {
				n++
			}
		}
		t.Error("Return value of the solution is not correct! Correct number is ", n)
	}
}
