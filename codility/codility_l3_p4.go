package codility

func Solve_l3_p4(S string, P []int, Q []int) []int {
	ps := make([][]int, 3)
	for i := 0; i < 3; i++ {
		ps[i] = make([]int, len(S)+1)
	}
	for i := 0; i < len(S); i++ {
		switch S[i] {
		case 'A':
			ps[0][i+1] += 1
		case 'C':
			ps[1][i+1] += 1
		case 'G':
			ps[2][i+1] += 1
		}
		for j := 0; j < 3; j++ {
			ps[j][i+1] += ps[j][i]
		}
	}

	res := make([]int, len(P))
	for i := 0; i < len(P); i++ {
		if ps[0][Q[i]+1]-ps[0][P[i]] > 0 {
			res[i] = 1
		} else if ps[1][Q[i]+1]-ps[1][P[i]] > 0 {
			res[i] = 2
		} else if ps[2][Q[i]+1]-ps[2][P[i]] > 0 {
			res[i] = 3
		} else {
			res[i] = 4
		}
	}

	return res
}
