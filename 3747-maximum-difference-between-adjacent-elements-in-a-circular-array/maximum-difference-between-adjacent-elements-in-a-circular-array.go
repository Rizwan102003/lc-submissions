func maxAdjacentDistance(nums []int) int {
    diff := 0
    n := len(nums)

    for i := 0; i < n; i++ {
        var curDiff int
        if i == n-1 {
            curDiff = nums[i] - nums[0]
        } else {
            curDiff = nums[i] - nums[i+1]
        }
        if abs(curDiff) > diff {
            diff = abs(curDiff)
        }
    }

    return diff
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}