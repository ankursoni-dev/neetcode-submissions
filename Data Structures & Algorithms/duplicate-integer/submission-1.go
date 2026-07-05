func hasDuplicate(nums []int) bool {
    if len(nums) <= 0 {
        return false
    }
    
    hash_map := make(map[int]int)

    i := 0
    for i = range(nums) {
        hash_map[nums[i]]++
        if hash_map[nums[i]] > 1 {
            return true
        }
    }
    if hash_map[nums[i]] > 1 {
        return true
    }
    return false
}
