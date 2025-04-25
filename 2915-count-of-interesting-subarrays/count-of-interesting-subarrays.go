// Complexity:
// Time O(N) and Space O(min(N,modulo)) where N is the length of nums.
func countInterestingSubarrays(nums []int, modulo int, k int) int64 {
	// We consider the number of indices where value % modulo equals k for
	// prefix arrays, and then further take the count modulo. Then, the problem
	// is similar to finding subarrays whose sum equal a specific value.
	prefixCnt := 0
	cntFreq := make(map[int]int)
	result := int64(0)

	cntFreq[0] = 1
	for _, num := range nums {
		if num%modulo == k {
			prefixCnt = (prefixCnt + 1) % modulo
		}
		if prefixCnt >= k {
			result += int64(cntFreq[prefixCnt-k])
		} else {
			result += int64(cntFreq[prefixCnt-k+modulo])
		}
		cntFreq[prefixCnt]++
	}
	return result
}