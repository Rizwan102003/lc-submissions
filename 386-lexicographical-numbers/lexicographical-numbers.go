func lexicalOrder(n int) []int {
    res := make([]int, 0, n)
    curr := 1
    for i := 0; i < n; i++ {
        res = append(res, curr)
        if curr*10 <= n {
            curr *= 10
        } else {
            if curr >= n {
                curr /= 10
            }
            curr += 1
            for curr%10 == 0 {
                curr /= 10
            }
        }
    }
    return res
}