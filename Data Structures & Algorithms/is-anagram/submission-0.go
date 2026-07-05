func isAnagram(s string, t string) bool {
	if len(s)!=len(t) {
		return false
	}

	hash_map1 := make(map[byte]int)
	hash_map2 := make(map[byte]int)

	for i := range(s) {
		hash_map1[s[i]]++
		hash_map2[t[i]]++
	}
	for i := range(s) {
		if hash_map1[s[i]] != hash_map2[s[i]]{
			return false
		}
	}
	return true
}
