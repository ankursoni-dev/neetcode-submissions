func groupAnagrams(strs []string) [][]string {
	hash_map := make(map[[26]int][]string)
	for _,str := range strs {
		count := [26]int{}
		for _, ch := range str {
			count[ch - 'a']++
		}
		hash_map[count] = append(hash_map[count],str);
	}

	res := [][]string{}
	for _, groups := range hash_map {
		res = append(res,groups)
	}
	return res
}
