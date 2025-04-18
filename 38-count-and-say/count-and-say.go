func countAndSay(n int) string {
	if n == 1 {
		return "1"
	} else {
		return say(countAndSay(n - 1))
	}
	return ""
}
func say(n string) string {
	result := []string{}
	tmp := n[0]
	count := 0
	for i := 0; i <= len(n)-1; i++ {
		if n[i] == tmp {
			count++
		} else {
			result = append(result, strconv.Itoa(count), string(tmp))
			tmp = n[i]
			count = 1
		}
	}
	result = append(result, strconv.Itoa(count), string(tmp))
	return strings.Join(result, "")
}