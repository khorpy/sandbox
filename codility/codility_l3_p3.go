package codility

func Solve_l3_p3(A []int) int {
	var slice_min, slice_avg float64

	ps := make([]int, len(A)+1)
	ps[1] = A[0]
	for i := 1; i < len(A); i++ {
		ps[i+1] = ps[i] + A[i]
	}

	i_min := 0
	slice_min = (float64(ps[2]) - float64(ps[0])) / 2.0
	for i := 0; i < len(A)-1; i++ {
		slice_avg = (float64(ps[i+2]) - float64(ps[i])) / 2.0
		if slice_avg < slice_min {
			i_min = i
			slice_min = slice_avg
		}
		if i < len(A)-2 {
			slice_avg = (float64(ps[i+3]) - float64(ps[i])) / 3.0
			if slice_avg < slice_min {
				i_min = i
				slice_min = slice_avg
			}
		}
	}

	return i_min
}
