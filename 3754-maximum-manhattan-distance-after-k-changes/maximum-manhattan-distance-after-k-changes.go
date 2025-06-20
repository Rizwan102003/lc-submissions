// Complexity:
// Time O(N) and Space O(1) where N is the length of s.
func maxDistance(s string, k int) int {
	netN := 0
	netE := 0
	result := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'N':
			netN++
		case 'S':
			netN--
		case 'E':
			netE++
		case 'W':
			netE--
		}

		result = max(result, min(i+1, abs(netN)+abs(netE)+2*k))
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}