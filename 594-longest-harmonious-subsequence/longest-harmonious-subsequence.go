func findLHS(nums []int) int {
    sort.Ints(nums)

    longest := 0
    l := 0
    for r := 0; r < len(nums); r++ {
        for nums[r]-nums[l] > 1 {
            l++
        }
        if nums[r]-nums[l] == 1 {
            longest = max(longest, r-l+1)
        }
    }
  
    return longest
}