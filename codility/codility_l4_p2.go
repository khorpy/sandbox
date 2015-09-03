package codility

func Solve_l4_p2(A []int) int {
	if len(A) < 3 {
		return 0
	}
	sorted := Merge_sort(A)
	for i := 0; i < len(sorted)-2; i++ {
		if sorted[i]+sorted[i+1] > sorted[i+2] && sorted[i+1]+sorted[i+2] > sorted[i] && sorted[i]+sorted[i+2] > sorted[i+1] {
			return 1
		}
	}
	return 0
}
