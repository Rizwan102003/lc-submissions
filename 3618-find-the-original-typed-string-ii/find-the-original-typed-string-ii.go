// Complexity:
// Time O(N+k^2) and Space O(N) where N is the length of word.
func possibleStringCount(word string, k int) int {
	if len(word) < k {
		return 0
	}
	if len(word) == k {
		return 1
	}

	modulo := 1_000_000_007

	// We solve an equivalent problem:
	//
	// Given an array of bags, each with bags[i] balls, what is the
	// number of ways to pick at least minBalls?

	bags := mapToBags(word)
	n := len(bags)
	minBalls := k - n

	totalWays := countTotalWays(bags, modulo)
	if minBalls <= 0 {
		return totalWays
	}

	invalidWays := countWays(bags, minBalls-1, modulo)
	return (totalWays - invalidWays + modulo) % modulo
}

// mapToBags returns an array containing len(part) - 1 for each part
// if we split the word when the character changes.
func mapToBags(word string) []int {
	result := []int{0}
	for i := 1; i < len(word); i++ {
		if word[i] != word[i-1] {
			result = append(result, 0)
		} else {
			result[len(result)-1]++
		}
	}
	return result
}

// countTotalWays returns the number of ways (mod the given modulo)
// to pick zero or more balls from bags.
func countTotalWays(bags []int, modulo int) int {
	result := int64(1)
	for _, num := range bags {
		result = (result * (1 + int64(num))) % int64(modulo)
	}
	return int(result)
}

// countWays returns the number of ways (mod the given modulo) to
// pick at most maxBalls from bags.
func countWays(bags []int, maxBalls int, modulo int) int {
	// dp[j]@i::= the number of ways to pick j balls from bags[0..=i]
	dp := make([]int64, maxBalls+1)
	dp[0] = 1 // base case i equals -1

	for i := range bags {
		wndSum := int64(0) // Sum of dp[j-bags[i]..<j]@i-1
		for j := max(0, (maxBalls+1)-bags[i]); j < (maxBalls + 1); j++ {
			wndSum = (wndSum + dp[j]) % int64(modulo)
		}

		for j := maxBalls; j >= 0; j-- {
			wndSum = (wndSum - dp[j] + int64(modulo)) % int64(modulo)
			if j >= bags[i] {
				wndSum = (wndSum + dp[j-bags[i]]) % int64(modulo)
			}
			dp[j] = (dp[j] + wndSum) % int64(modulo)
		}
	}

	result := 0
	for _, count := range dp {
		result = (result + int(count)) % modulo
	}
	return result
}