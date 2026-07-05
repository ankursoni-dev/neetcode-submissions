func twoSum(nums []int, target int) []int {
    hash_map := make(map[int]int)
	for i:= range nums {
		diff := target - nums[i];
		_, ok := hash_map[diff]
		if ok {
			return []int{hash_map[diff],i}
		}
		hash_map[nums[i]] = i;
	}
	return []int{0,0}
}
