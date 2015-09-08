package codility

func Solve_l6_p1(A []int) int {
	if len(A) == 0 {
		return -1
	}
	i, v := 0, 0
	for j := 0; j < len(A); j++ {
		if i == 0 {
			i++
			v = A[j]
		} else {
			if v == A[j] {
				i++
			} else {
				i--
			}
		}
	}
	if i > 0 {
		indx := -1
		count := 0
		for j := 0; j < len(A); j++ {
			if A[j] == v {
				count++
				if indx == -1 {
					indx = j
				}
			}
		}
		if count > len(A)/2 {
			return indx
		} else {
			return -1
		}
	} else {
		return -1
	}
}
