package codility

func Solve_l6_p2(A []int) int {
	if len(A) < 2 {
		return 0
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
		count := 0
		for j := 0; j < len(A); j++ {
			if A[j] == v {
				count++
			}
		}
		if count > len(A)/2 {
			eq := 0
			count_left := 0
			for j := 0; j < len(A)-1; j++ {
				if A[j] == v {
					count_left++
				}
				if count_left > (j+1)/2 && count-count_left > (len(A)-j-1)/2 {
					eq++
				}
			}
			return eq
		} else {
			return 0
		}
	} else {
		return 0
	}
}
