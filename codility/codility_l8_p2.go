package codility

import "math"

func Solve_l8_p2(N int) int {
	s := int(math.Sqrt(float64(N)))
	for N%s != 0 {
		s--
	}
	return 2 * (s + N/s)
}
