func possibleStringCount(word string) int {
    cnt := 1
    for i := 1; i < len(word); i++ {
        if word[i] == word[i-1] {
            cnt++
        }
    }
    return cnt
}