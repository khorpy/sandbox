package codility

import "testing"

func TestSolve_l3_p4(t *testing.T) {
	S := "CAGCCTA"
	P := []int{2, 5, 0}
	Q := []int{4, 5, 6}
	exp := []int{2, 4, 1}
	res := Solve_l3_p4(S, P, Q)

	for i := 0; i < len(res); i++ {
		if exp[i] != res[i] {
			t.Error("Return value of the solution is not correct!", res)
			break
		}
	}

	S = "A"
	P = []int{0}
	Q = []int{0}
	exp = []int{1}
	res = Solve_l3_p4(S, P, Q)

	for i := 0; i < len(res); i++ {
		if exp[i] != res[i] {
			t.Error("Return value of the solution is not correct!", res)
			break
		}
	}
}
