import (
    "fmt"
    // "sort"
)

func partitionArray(nums []int, k int) int {
    // 1. O(nlogn), O(1)
    // sort.Slice(nums, func(i, j int) bool {
    //     return nums[i] < nums[j]
    // })
    // start := - (k + 1)  // since nums[0] should start with new group
    // res := 0
    // for _, val := range nums {
    //     if val - start > k {
    //         start = val
    //         res++
    //     }
    // }
    // return res

    // 2. O(n), O(n) Counting sort and then same subarray counting
    maxval := nums[0]
    for _, val := range nums {
        maxval = max(maxval, val)
    }
    freq := make([]int, maxval + 1)
    for _, val := range nums {
        freq[val]++
    }
    res := 0
    start := -(k + 1)
    for val := 0; val <= maxval; val++ {
        if val - start > k  && freq[val] != 0 {
            start = val
            res++
        }
    }
    return res
}