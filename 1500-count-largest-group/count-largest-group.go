func countLargestGroup(n int) int {
	mapGroup := make(map[int]int)
	for i := 1; i <= n; i++ {
		sod := sumOfDigit(i)
		mapGroup[sod]++
	}

	var result, tmpMax int
	for _, val := range mapGroup {
		if tmpMax < val {
			result = 0
			tmpMax = val
		}

		if val == tmpMax {
			result++
		}
	}

	return result
}

func sumOfDigit(n int) int {
	var sum int
	for n >= 1 {
		sum += (n % 10)
		n /= 10
	}
	return sum
}