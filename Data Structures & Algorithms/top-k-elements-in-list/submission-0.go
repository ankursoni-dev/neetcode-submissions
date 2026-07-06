func topKFrequent(nums []int, k int) []int {
	hash_map := make(map[int]int)
	for _,val := range nums {
		hash_map[val]++
	}
	freq_map := make([][]int, len(nums)+1);
	for num,cnt := range hash_map {
		freq_map[cnt] = append(freq_map[cnt], num)
	}
	answ_arr := []int{}
	for i:= len(freq_map) - 1; i >0 ; i-- {
		for _,val := range freq_map[i]{
			answ_arr = append(answ_arr,val)
			if len(answ_arr) == k {
				return answ_arr
			}
		}
	}
	return answ_arr
}
