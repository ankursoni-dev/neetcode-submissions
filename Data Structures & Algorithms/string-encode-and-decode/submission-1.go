type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	encoded_string := "";
	for _,str := range strs {
		num_to_str := strconv.Itoa(len(str))
		encoded_string += num_to_str
		encoded_string += "#"
		encoded_string += str
	}
	return encoded_string
}

func (s *Solution) Decode(encoded string) []string {
	decoded_array := []string{}
	i := 0
	for i < len(encoded) {
		j := i
		for encoded[j] != '#' {
			j++
		}
		length,_ := strconv.Atoi(encoded[i:j])
		i = j + 1
		decoded_array = append(decoded_array,encoded[i:i+length])
		i += length
	}
	return decoded_array
}
