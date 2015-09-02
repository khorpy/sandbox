package codility

func Solve_l4_p1(A []int) int {
	A_sorted := Merge_sort(A)
	v1 := A_sorted[len(A)-1] * A_sorted[len(A)-2] * A_sorted[len(A)-3]
	v2 := A_sorted[0] * A_sorted[1] * A_sorted[len(A)-1]

	if v1 > v2 {
		return v1
	} else {
		return v2
	}
}

func Merge_sort(a []int) []int {
	if len(a) < 2 {
		return a
	} else {
		return merge_sorted(Merge_sort(a[:len(a)/2]), Merge_sort(a[len(a)/2:]))
	}
}

func merge_sorted(a, b []int) []int {
	merged := make([]int, len(a)+len(b))
	ia, ib := 0, 0
	for ia+ib < len(merged) {
		if ia < len(a) && ib < len(b) {
			if a[ia] < b[ib] {
				merged[ia+ib] = a[ia]
				ia++
			} else {
				merged[ia+ib] = b[ib]
				ib++
			}
		} else if ia < len(a) {
			merged[ia+ib] = a[ia]
			ia++
		} else {
			merged[ia+ib] = b[ib]
			ib++
		}
	}
	return merged
}
