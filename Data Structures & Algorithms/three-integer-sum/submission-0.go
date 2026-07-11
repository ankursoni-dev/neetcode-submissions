import "slices"
func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	ans := [][]int{}
	ans_map := make(map[[3]int]int)
	for ind,num := range nums {
		ptr1 := ind+1;
		ptr2 := len(nums) - 1;
		for ptr1 < ptr2 {
			sum := num + nums[ptr1] + nums[ptr2]
			if sum > 0 {
				ptr2--
			} else if sum < 0 {
				ptr1++
			} else {
				one_ans := [3]int{num,nums[ptr1], nums[ptr2]}
				if _,ok:= ans_map[one_ans]; !ok {
					ans_map[one_ans]++;
					ans = append(ans, []int{num,nums[ptr1], nums[ptr2]})
				}
				ptr2--
				ptr1++
			}
		}
	}
	return ans
}
// [-1,0,-1,2,-1,-4]
// [-4,-1,-1,-1,0,2]