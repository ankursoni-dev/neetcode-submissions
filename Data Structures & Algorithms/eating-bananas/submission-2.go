func minEatingSpeed(piles []int, h int) int {
	lPace,rPace:=1,maxElOfArray(piles)
	resPace := 1;
	for lPace <= rPace {
		pace := lPace + (rPace - lPace)/2
		time := calcPileConsumptionTime(piles,pace)
		if time <= h {
			resPace = pace
			rPace = pace - 1
		} else {
			lPace = pace + 1
		}
	}
	return resPace
}

func calcPileConsumptionTime (piles []int, pace int) int {
	time := 0
	for _,pile := range piles {
		if pile%pace == 0 {
			time += pile/pace
		}else {
			time += pile/pace + 1
		}
	}
	return time
}

func maxElOfArray(array []int) int {
	max := math.MinInt
	for _,el := range array {
		if el > max {
			max = el
		}
	}
	return max
}