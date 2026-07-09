type CustomString string
func (c CustomString) ToLower() CustomString {
    b := []byte(c)
    for i, ch := range b {
        if ch >= 'A' && ch <= 'Z' {
            b[i] = ch + 32
        }
    }
    return CustomString(b)
}

func IsAlphaNumeric(r byte) bool {
    return (r >= '0' && r <= '9') ||
           (r >= 'A' && r <= 'Z') ||
           (r >= 'a' && r <= 'z')
}

func isPalindrome(s string) bool {
	str := CustomString(s)
	str = str.ToLower()
	i, j := 0, len(s)-1
	for i <= j{
		for !IsAlphaNumeric(str[j]) && j > 0{j--}
		for !IsAlphaNumeric(str[i]) && i < len(s) - 1 {i++}
		if IsAlphaNumeric(str[i]) && str[i] != str[j]{
			return false
		}
		i++;
		j--;
	}
	return true
}