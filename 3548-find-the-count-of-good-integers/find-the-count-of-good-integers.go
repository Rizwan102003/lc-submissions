func countGoodIntegers(n int, k int) int64 {
    const mod = 1e9 + 7

    // Precompute factorials up to n for permutation calculations
    factorial := make([]int64, n+1)
    factorial[0] = 1
    for i := 1; i <= n; i++ {
        factorial[i] = (factorial[i-1] * int64(i)) % mod
    }

    // Generate all unique digit multisets that can form k-palindromic numbers
    digitMultisets := make(map[string]bool)

    // The smallest n-digit number (10^(n-1))
    min := int(math.Pow10(n - 1))
    // The largest n-digit number (10^n - 1)
    max := int(math.Pow10(n)) - 1

    // Generate palindromes efficiently by constructing first half
    halfLen := (n + 1) / 2
    start := int(math.Pow10(halfLen - 1))
    end := int(math.Pow10(halfLen)) - 1

    for firstHalf := start; firstHalf <= end; firstHalf++ {
        // Construct the full palindrome
        s := strconv.Itoa(firstHalf)
        var palindromeStr string
        if n%2 == 0 {
            palindromeStr = s + reverseString(s)
        } else {
            palindromeStr = s + reverseString(s)[1:]
        }

        // Convert to integer and check divisibility by k
        palindrome, _ := strconv.Atoi(palindromeStr)
        if palindrome < min || palindrome > max {
            continue // Skip numbers outside our n-digit range
        }

        if palindrome%k == 0 {
            // Create sorted digit representation as multiset key
            digits := strings.Split(palindromeStr, "")
            sort.Strings(digits)
            multisetKey := strings.Join(digits, "")
            digitMultisets[multisetKey] = true
        }
    }

    var count int64 = 0

    // For each unique digit multiset, calculate valid permutations
    for multiset := range digitMultisets {
        digitCounts := make([]int, 10)
        for _, ch := range multiset {
            digitCounts[ch-'0']++
        }

        // Calculate total permutations (n! / (count_d0! * count_d1! * ...))
        permutations := factorial[n]
        for _, cnt := range digitCounts {
            if cnt > 1 {
                permutations = permutations * modInverse(factorial[cnt], mod) % mod
            }
        }

        // Subtract permutations with leading zeros
        if digitCounts[0] > 0 {
            invalidPerms := factorial[n-1] // First digit is fixed as 0
            for d, cnt := range digitCounts {
                if d == 0 {
                    invalidPerms = invalidPerms * modInverse(factorial[cnt-1], mod) % mod
                } else {
                    invalidPerms = invalidPerms * modInverse(factorial[cnt], mod) % mod
                }
            }
            permutations = (permutations - invalidPerms + mod) % mod
        }

        count = (count + permutations) % mod
    }

    return count
}

// Helper function to compute modular inverse using Fermat's Little Theorem
func modInverse(a int64, mod int64) int64 {
    return powMod(a, mod-2, mod)
}

// Helper function for modular exponentiation
func powMod(a, b, mod int64) int64 {
    var result int64 = 1
    a = a % mod
    for b > 0 {
        if b%2 == 1 {
            result = (result * a) % mod
        }
        a = (a * a) % mod
        b = b / 2
    }
    return result
}

// Helper function to reverse a string
func reverseString(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}