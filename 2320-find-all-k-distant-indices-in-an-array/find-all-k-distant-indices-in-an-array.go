func findKDistantIndices(nums []int, key int, k int) []int {
    m := make(map[int]int)
    ans := []int{}

    l := 0
    start := 0
    for i:=0 ; i<=k && i<len(nums) ; i++ {
        if nums[i] == key {
            m[i]++
        }
        l++
    }

    if len(m) > 0 {
        ans = append(ans , 0)
    }

    for i:=1 ; i<len(nums) ; i++ {
        if i+k < len(nums) && nums[i+k] == key {
            m[i+k]++
        }
        l++

        if l > 2*k + 1 {
            if nums[start] == key {
                delete(m , start)
            }
            start++
            l--
        }

        if len(m) > 0 {
            ans = append(ans , i)
        }

    }

    return ans
}