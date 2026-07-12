func trap(height []int) int {
	if len(height) == 0{
		return 0
	}

	l,r:=0,len(height)-1;
	lMax,rMax:=height[l],height[r]
	res:=0;

	for l < r{
		if lMax < rMax {
			l++;
			lMax = max(lMax,height[l])
			res += lMax - height[l]
		} else {
			r--;
			rMax = max(rMax,height[r])
			res += rMax - height[r]
		}
	}
	return res
}
