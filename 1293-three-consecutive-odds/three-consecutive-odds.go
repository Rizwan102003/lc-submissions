func threeConsecutiveOdds(arr []int) bool {
    var window int
    for i := 0; i < 3 && i < len(arr); i++ {
        window += arr[i] % 2
    }
    if window == 3 {
        return true
    }
    for i := 3; i < len(arr); i++ {
        window += arr[i] % 2
        window -= arr[i-3] % 2
        if window == 3 {
            return true
        }
    }
    return false
}