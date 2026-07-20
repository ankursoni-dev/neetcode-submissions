func maxProfit(prices []int) int {
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}
	if len(prices) == 2 && prices[0] > prices[1] {
		return 0
	}
	max_profit := math.MinInt
	for i:=0; i<len(prices); i++ {
		j := i+1
		curr_profit := 0
		fmt.Println("before : ", j, curr_profit)
		for j < len(prices) && prices[j] > prices[i] {
			curr_profit = prices[j] - prices[i]
			max_profit = max(max_profit,curr_profit)
			j++;
			fmt.Println("during : ", j, curr_profit)
		}
		fmt.Println("after : ", j, curr_profit)
		max_profit = max(max_profit,curr_profit)
		i = j - 1
	}
	return max_profit
}
