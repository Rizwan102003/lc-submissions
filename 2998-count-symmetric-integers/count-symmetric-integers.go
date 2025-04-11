func countSymmetricIntegers(low int, high int) int {
 var count int
    for i := low; i <= high; i++ {
        // Count the number of digits in i
        n := i
        numDigits := 0
        for n > 0 {
            n /= 10
            numDigits++
        }
        // Skip odd-digit numbers
        if numDigits%2 == 1 {
            continue
        }
        // Calculate the sum of the first half of digits and the sum of the second half
        n = i
        sum1 := 0
        sum2 := 0
        for j := 0; j < numDigits; j++ {
            digit := n % 10
            n /= 10
            if j < numDigits/2 {
                sum1 += digit
            } else {
                sum2 += digit
            }
        }
        // Check if the sums are equal
        if sum1 == sum2 {
            count++
        }
    }
    return count
}