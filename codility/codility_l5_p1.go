package codility

func Solve_l5_p1(S string) int {
	counter := 0
	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			counter++
		} else {
			counter--
		}
		if counter < 0 {
			return 0
		}
	}

	if counter > 0 {
		return 0
	} else {
		return 1
	}
}
