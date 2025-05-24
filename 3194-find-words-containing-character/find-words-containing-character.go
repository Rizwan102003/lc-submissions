// Complexity:
// Time O(L) and Space O(N) where N and L are the length and the flattened
// length of words, respectively.
func findWordsContaining(words []string, x byte) []int {
	result := make([]int, 0)
	for i, word := range words {
		if strings.IndexByte(word, x) != -1 {
			result = append(result, i)
		}
	}
	return result
}