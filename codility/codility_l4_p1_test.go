package codility

import "testing"

func TestSolve_l4_p1(t *testing.T) {
	A := []int{-3, 1, 2, -2, 5, 6}

	if Solve_l4_p1(A) != 60 {
		t.Error("Return value of the solution is not correct!")
	}

	A = []int{-5, 5, -5, 4}

	if Solve_l4_p1(A) != 125 {
		t.Error("Return value of the solution is not correct!")
	}

}

func TestMerge_sort(t *testing.T) {
	input := []int{-10, 2, 8, 0, 13, 12}
	output_exp := []int{-10, 0, 2, 8, 12, 13}

	output := Merge_sort(input)

	for i := 0; i < len(input); i++ {
		if output[i] != output_exp[i] {
			t.Error("Array is not sorted: ", output)
		}
	}
}
