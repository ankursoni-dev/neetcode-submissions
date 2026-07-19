func search(nums []int, target int) int {
	lowIdx,highIdx:=0,len(nums)-1
	for lowIdx <= highIdx {
		midIdx := lowIdx + (highIdx-lowIdx)/2
		lowNum,midNum,highNum := nums[lowIdx],nums[midIdx],nums[highIdx]
		if midNum == target {
			return midIdx
		} else if lowNum <= midNum {
			if lowNum <= target && midNum > target {
				highIdx = midIdx - 1
			} else {
				lowIdx = midIdx + 1
			}
		} else {
			if midNum < target && highNum >= target {
				lowIdx = midIdx + 1
			} else {
				highIdx = midIdx - 1
			}
		}
	}
	return -1
}
