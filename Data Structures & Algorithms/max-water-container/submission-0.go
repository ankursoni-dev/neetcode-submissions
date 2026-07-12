func maxArea(heights []int) int {
	max_vol := 0
	lptr,rptr:= 0,len(heights)-1;
	for lptr < rptr {
		min_height := min(heights[lptr],heights[rptr])
		vol := min_height*(rptr-lptr)
		max_vol = max(max_vol,vol)
		if heights[lptr] < heights[rptr] {
			lptr++
		}else {
			rptr--
		}
	}
	return max_vol
}
