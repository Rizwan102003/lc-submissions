func divideString(s string, k int, fill byte) []string {
	result := []string{}
	n := len(s)

	for i := 0; i < n; i += k {
		end := i + k
		if end > n {
			chunk := s[i:]
			padding := make([]byte, k-len(chunk))
			for j := range padding {
				padding[j] = fill
			}
			chunk += string(padding)
			result = append(result, chunk)
		} else {
			result = append(result, s[i:end])
		}
	}

	return result
}