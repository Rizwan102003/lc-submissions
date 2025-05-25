func longestPalindrome(words []string) int {
    wordMap := map[string]int{}
    for i := 0 ; i < len(words); i++ {
        wordMap[words[i]]++
    }
    
    output := 0
    // find pairs
    for word, wordCount := range wordMap {
        reversed := reverseString(word)
        if reversed == word {
            output += len(word) * (wordCount/2) * 2
            wordMap[word] = wordMap[word] % 2
            if wordMap[word] == 0 {
                delete(wordMap, word)
            }
        } else {
            reverseCount := wordMap[reversed]
            minCount := min(wordCount,reverseCount)
            output += minCount * len(word) * 2
            wordMap[word] -= minCount
            wordMap[reversed] -= minCount
            if wordMap[word] == 0 {
                delete(wordMap, word)
            }
            if wordMap[reversed] == 0 {
                delete(wordMap, reversed)
            }
        }
    }
    
    // find longest palindrome
    longest := ""
    for word, _ := range wordMap {
        if isPalindrome(word) {
            if len(word) > len(longest) {
                longest = word
            }
        }
    }

    return output + len(longest)
}

func min(a,b int) int{
    if a < b {
        return a
    }
    return b
}

func reverseString(s string) string {
    runes := []rune(s)        
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i] 
    }
    return string(runes)
}

func isPalindrome(word string) bool {
    left := 0
    right := len(word)-1
    for left < right {
        if word[left] != word[right]{
            return false
        }
        left++
        right--
    }   
    return true 
}