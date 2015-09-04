package codility

func Solve_l5_p3(S string) int {
	bs := make([]uint8, 0)
	for i := 0; i < len(S); i++ {
		switch S[i] {
		case '(', '[', '{':
			bs = append(bs, S[i])
		case ')', ']', '}':
			if len(bs) == 0 {
				return 0
			}
			if S[i] == ')' && bs[len(bs)-1] == '(' || S[i] == ']' && bs[len(bs)-1] == '[' ||
				S[i] == '}' && bs[len(bs)-1] == '{' {
				bs = bs[:len(bs)-1]
			} else {
				return 0
			}
		}
	}
	if len(bs) > 0 {
		return 0
	} else {
		return 1
	}
}
