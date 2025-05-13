func lengthAfterTransformations(s string, t int) int {
     const MOD = 1000000007

	count := make([]int64, 26)
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
	}

	for i := 0; i < t; i++ {
		newCount := make([]int64, 26)

		for j := 0; j < 25; j++ {
			newCount[j+1] = (newCount[j+1] + count[j]) % MOD
		}

		newCount[0] = (newCount[0] + count[25]) % MOD
		newCount[1] = (newCount[1] + count[25]) % MOD 

		count = newCount
	}

	var totalLength int64
	for i := 0; i < 26; i++ {
		totalLength = (totalLength + count[i]) % MOD
	}

	return int(totalLength)
}