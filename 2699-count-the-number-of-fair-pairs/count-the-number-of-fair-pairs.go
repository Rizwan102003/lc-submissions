func countFairPairs(nums []int, lower int, upper int) int64 {
    slices.Sort(nums)
    return countLess(nums, upper) - countLess(nums, lower - 1)
}
func countLess(nums []int, sum int) int64 {
    result := int64(0)
    for i, j := 0, len(nums) - 1; i < j; i++ {
        for i < j && nums[i] + nums[j] > sum {
            j--
        }
        result += int64(j - i)
    }
    return result;
}