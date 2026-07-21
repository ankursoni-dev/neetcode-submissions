func lengthOfLongestSubstring(s string) int {
	set := make(map[byte]int)
	l,max_len := 0,0
	for ptr := 0; ptr < len(s); ptr++ {
		if idx, ok := set[s[ptr]]; ok {
			l = max(idx+1,l)
		}
		set[s[ptr]] = ptr
		if ptr - l + 1 > max_len {
			max_len = ptr - l + 1
		}
	}
	return max_len
}
