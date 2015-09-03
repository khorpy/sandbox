package codility

func Solve_l4_p3(A []int) int {
	if len(A) < 2 {
		return len(A)
	}
	sorted_a := Merge_sort(A)
	counter := 1
	for i := 1; i < len(A); i++ {
		if sorted_a[i] != sorted_a[i-1] {
			counter++
		}
	}

	return counter
}
