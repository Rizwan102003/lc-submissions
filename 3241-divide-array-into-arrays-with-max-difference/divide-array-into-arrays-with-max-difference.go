func divideArray(nums []int, k int) [][]int {

    sort.Ints(nums)
    three_consecutive_elements := [][]int{}
    n := len(nums)
    
    for i := 0; i < n;{
        group := nums[i:i+3]
        if len(group) == 3 && group[2] - group[0] <= k {
            three_consecutive_elements = append(three_consecutive_elements, group)
            i += 3
        } else {
            return [][]int{}
        }
    
    }
    return three_consecutive_elements

    
}