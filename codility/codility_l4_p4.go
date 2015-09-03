package codility

func Solve_l4_p4(A []int) int {
	left := make([]int, len(A))
	right := make([]int, len(A))

	for i := 0; i < len(A); i++ {
		left[i] = i - A[i]
		right[i] = i + A[i]
	}

	left = Merge_sort(left)
	right = Merge_sort(right)

	result, counter, li, ri := 0, 0, 0, 0

	for li < len(left) || ri < len(right) {
		if li < len(left) && ri < len(right) {
			if left[li] <= right[ri] {
				result += counter
				if result > 10000000 {
					return -1
				}
				counter++
				li++
			} else {
				counter--
				ri++
			}
		} else if li < len(left) {
			result += counter
			if result > 10000000 {
				return -1
			}
			counter++
			li++
		} else {
			counter--
			ri++
		}
	}
	return result
}
