func find(char byte, representative *[26]byte) byte {
    if (*representative)[char] == char {
        return char
    }
    
    (*representative)[char] = find(representative[char], representative)
    
    return (*representative)[char]
}

func performUnion(x, y byte, representative *[26]byte) {
    x, y = find(x, representative), find(y, representative)

    switch {
    case x < y:
        (*representative)[y] = x
    case x > y:
        (*representative)[x] = y
    }
}

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
    representative := [26]byte{}
    for i := byte(0); i < 26; i++ {
        representative[i] = i
    }

    l1, l2 := []byte(s1), []byte(s2)
    for i := range l1 {
        performUnion(l1[i] - 'a', l2[i] - 'a', &representative)
    }

    result := []byte{}
    for _, c := range []byte(baseStr) {
        result = append(result, find(c - 'a', &representative) + 'a')
    }

    return string(result)
}