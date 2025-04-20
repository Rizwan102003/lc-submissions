func numRabbits(answers []int) int {
    m := make(map[int]int)
	for _, a := range answers {
		m[a+1]++ // a+1 since self included
	}

	res := 0
	for k, v := range m {
		res += k * (((v - 1) / k) + 1)
	}
	return res
}