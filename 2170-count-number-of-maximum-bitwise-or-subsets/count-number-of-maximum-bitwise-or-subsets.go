func countMaxOrSubsets(nums []int) int {
    subset := subsetsBit(nums)
    hash := map[int]int{}
    for _, h := range subset {
        hash[Or(h)]++
    }

    max := 0
    for k, _ := range hash {
        if max < k {
            max = k
        }
    }

    return hash[max]
}

func Or(arr []int) int{
    if len(arr) == 0 {
        return 0
    }

    ans := arr[0]
    for i := 1; i < len(arr); i++ {
        ans |= arr[i]
    }

    return ans
}

func subsetsBit(nums []int) [][]int {
    n := len(nums)
    res := make([][]int, 0, n)

    for x := 0; x < 1 << n; x++ {
        ele := make([]int, 0, n)
        for b:=0; b<n; b++ {
            if x & (1 << b) > 0 {
                ele = append(ele, nums[b])
            }
        }

        res = append(res, ele)
    }

    return res
}