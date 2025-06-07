func clearStars(s string) string {
    occToIndices := [26][]int{} // indices where each char occured, when removing always remove last
    deleted := map[int]bool{} // to mark when we delete from an index to make life easier
    for i := range s {
        if s[i] != '*' {
            occToIndices[int(s[i]-'a')] = append(occToIndices[int(s[i]-'a')], i)
            continue
        }

        // do O(26) here to check the smallest char and remove it
        for j := 0; j < 26 ; j++ {
            if len(occToIndices[j]) > 0 {
                toDelete :=  occToIndices[j][len(occToIndices[j])-1]
                occToIndices[j] = occToIndices[j][:len(occToIndices[j])-1]
                deleted[toDelete] = true
                break
            }
        }
    }

    var res strings.Builder
    for i := range s {
        if !deleted[i] && s[i] != '*'{
            res.WriteByte(s[i])
        }
    }
    return res.String()
}