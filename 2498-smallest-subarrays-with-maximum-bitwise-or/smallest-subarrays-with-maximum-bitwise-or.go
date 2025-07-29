func smallestSubarrays(nums []int) []int {
    leftMostIdx, res := [32]int{}, make([]int, len(nums))
    for i := len(nums) - 1; i >= 0; i--{
        endIdx := i
        for bitIdx := 0; bitIdx < 32; bitIdx++{
            mask := 1 << bitIdx
            if nums[i]  & mask > 0 {
                leftMostIdx[bitIdx] = i
            }
            endIdx = max(endIdx, leftMostIdx[bitIdx])
        }
        res[i] = endIdx - i + 1
    }
    return res
}