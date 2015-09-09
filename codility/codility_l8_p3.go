package codility

import "math"

func Solve_l8_p3(A []int) int {
	if len(A) < 3 {
		return 0
	}
	peaks := make([]int, len(A))
	for i := 1; i < len(A)-1; i++ {
		if A[i] > A[i-1] && A[i] > A[i+1] {
			peaks[i] = 1
		}
	}

	pp := make([]int, len(A)+1)
	for i := 0; i < len(peaks); i++ {
		pp[i+1] = pp[i] + peaks[i]
	}

	imax := 0
	if pp[len(A)] > 0 {
		imax = 1
	} else {
		return 0
	}

	for i := int(math.Sqrt(float64(len(A)))); i > 1; i-- {
		if len(A)%i == 0 {
			check := 1
			for j := 1; i*j <= len(A); j++ {
				if pp[i*j]-pp[i*j-i] == 0 {
					check = 0
					break
				}
			}
			if check == 1 && len(A)/i > imax {
				imax = len(A) / i
			}
			k := len(A) / i
			if i != k {
				check = 1
				for j := 1; k*j <= len(A); j++ {
					if pp[k*j]-pp[k*j-k] == 0 {
						check = 0
						break
					}
				}
				if check == 1 && len(A)/k > imax {
					imax = len(A) / k
				}
			}
		}
	}

	return imax
}
