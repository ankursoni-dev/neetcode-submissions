func productExceptSelf(nums []int) []int {
	suffix_arr := []int{1}
	prefix_arr := []int{1}

	for i := 1; i< len(nums); i++ {
		prefix_arr = append(prefix_arr, prefix_arr[i-1]*nums[i-1])
		last_idx := len(nums) - i
		suffix_arr = append(suffix_arr, suffix_arr[i - 1]*nums[last_idx])
	}

	result_arr := []int{}
	for idx,num := range prefix_arr {
		result_arr = append(result_arr,num*suffix_arr[len(nums)-idx-1])
	}
	return result_arr
}
// 1 1 2 8
// 1 6 24 48

// 1 -1 0 0 0
// 1 3 6 6 0