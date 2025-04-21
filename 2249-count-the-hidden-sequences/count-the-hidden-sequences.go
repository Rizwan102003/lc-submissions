func numberOfArrays(differences []int, lower int, upper int) int {
    x, y, cur := 0, 0, 0
    for _, d := range differences {
        cur += d
        if cur < x {
            x = cur
        }
        if cur > y {
            y = cur
        }
    }
    res := (upper - lower) - (y - x) + 1
    if res < 0 {
        return 0
    }
    return res
}