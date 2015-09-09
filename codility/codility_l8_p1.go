package codility

import "math"

func Solve_l8_p1(N int) int {
	if N < 3 {
		return N
	}
	count := 2
	for i := int(math.Sqrt(float64(N))); i > 1; i-- {
		if N%i == 0 {
			count += 2
		}
		if i*i == N {
			count--
		}
	}
	return count
}
