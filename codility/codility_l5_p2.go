package codility

func Solve_l5_p2(H []int) int {
	height := 0
	blocks := make([]int, 0)
	counter := 0
	for i := 0; i < len(H); i++ {
		if H[i] > height {
			blocks = append(blocks, H[i]-height)
			height = H[i]
			counter++
		} else if H[i] < height {
			for height > H[i] {
				height -= blocks[len(blocks)-1]
				blocks = blocks[:len(blocks)-1]
			}
			if H[i] != height {
				blocks = append(blocks, H[i]-height)
				height = H[i]
				counter++
			}
		}
	}
	return counter
}
