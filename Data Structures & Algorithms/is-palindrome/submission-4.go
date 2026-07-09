func isPalindrome(s string) bool {
	str := strings.ToLower(s)
	i, j := 0, len(s)-1
	for i < j{
		for i<j && !unicode.IsLetter(rune(str[i])) && !unicode.IsNumber(rune(str[i])) {i++}
		for i<j && !unicode.IsLetter(rune(str[j])) && !unicode.IsNumber(rune(str[j])) {j--}
		if str[i] != str[j]{
			return false
		}
		i++;
		j--;
	}
	return true
}