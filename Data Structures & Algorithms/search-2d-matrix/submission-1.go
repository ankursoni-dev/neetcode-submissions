func isPresent(array []int, target int) bool {
	l,r := 0, len(array)
	for l <=r {
		mid := l + (r-l)/2
		if array[mid] == target {
			return true
		}else if array[mid] > target {
			r = mid - 1
		}else {
			l = mid + 1
		}
	}
	return false
}

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix[0])
	n := len(matrix)
	i := 0
	for i < n {
		if target == matrix[i][m-1] || target == matrix[i][0] {
			return true
		}else if target < matrix[i][m-1] {
			return isPresent(matrix[i], target)
		}
		i++
	}
	return false
}

