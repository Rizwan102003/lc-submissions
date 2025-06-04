func answerString(w string, n int) string {
    if n == 1 {
        return w
    }
    lenAns := len(w) - n + 1
    ans := ""
    for i := 0; i < len(w); i++ {
        ans = max(ans, w[i:min(i + lenAns, len(w))])
    }
    return ans
}