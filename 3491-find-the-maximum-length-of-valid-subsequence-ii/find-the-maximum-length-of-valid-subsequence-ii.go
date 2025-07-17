package main

import (
	"fmt"
)

func maximumLength(nums []int, k int) int {
	maxLength := 2

	for targetMod := 0; targetMod < k; targetMod++ {
		remainderDP := make([]int, k)

		for _, num := range nums {
			numMod := num % k
			requiredMod := (targetMod - numMod + k) % k
			remainderDP[numMod] = remainderDP[requiredMod] + 1
		}

		for _, length := range remainderDP {
			if length > maxLength {
				maxLength = length
			}
		}
	}

	return maxLength
}