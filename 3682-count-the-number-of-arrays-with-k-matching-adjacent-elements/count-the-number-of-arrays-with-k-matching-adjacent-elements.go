const MOD = 1000000007 
func countGoodArrays(n int, m int, k int) int {
    fact := make([]int64, n)
    fact[0] = 1
    for i := 1; i < n; i++ {
        fact[i] = (fact[i-1] * int64(i)) % MOD
    }
    pow := func(base, exp, mod int64) int64 {
        result := int64(1)
        base %= mod
        for exp > 0 {
            if exp%2 == 1 {
                result = (result * base) % mod
            }
            base = (base * base) % mod
            exp /= 2 
        }
        return result
    }
    modInverse := func(a int64) int64 {
        return pow(a, MOD-2, MOD)//a/b=a*b^-1 Mod --fermats little theorem
    }
    combination := func(n, r int) int64 {
        numerator := fact[n-1]//--(n-1)!/(n-1-r)!
        denominator := (fact[r] * fact[n-1-r]) % MOD
        return (numerator * modInverse(denominator)) % MOD
    }
    ways := combination(n, k)//k positions in n for maximum m
    result := int64(m)//only one choice m
    choices := pow(int64(m-1), int64(n-1-k), MOD)//remaining choices
    
    result = (result * ways) % MOD//total number of ways to position max m exactly k number of times
    result = (result * choices) % MOD//total number of choices <m for remaining pos
    
    return int(result)
}