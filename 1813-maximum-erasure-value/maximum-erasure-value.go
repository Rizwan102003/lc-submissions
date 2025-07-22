// Complexity:
// Time O(N) and Space O(N) where N is the length of nums.
func maximumUniqueSubarray(nums []int) int {
	result := nums[0]
	start := 0
	windowSum := 0
	inWindow := make(map[int]bool)
	for end := range nums {
		for inWindow[nums[end]] {
			inWindow[nums[start]] = false
			windowSum -= nums[start]
			start++
		}
		windowSum += nums[end]
		inWindow[nums[end]] = true
		result = max(result, windowSum)
	}
	return result
}