func isValidSudoku(board [][]byte) bool {
	for row:=0;row<9;row++{
		row_map:=make(map[byte]bool)
		for i:=0;i<9;i++{
			if board[row][i] == '.' {
				continue
			}
			num:=board[row][i]
			if row_map[num] {
				return false
			}
			row_map[num]=true
		}
	}
	for col:=0;col<9;col++{
		col_map:=make(map[byte]bool)
		for i:=0;i<9;i++{
			if board[i][col] == '.' {
				continue
			}
			num:=board[i][col]
			if col_map[num] {
				return false
			}
			col_map[num]=true
		}
	}
	for sqr:=0;sqr<9;sqr++{
		sqr_map:=make(map[byte]bool)
		for i:=0;i<3;i++{
			for j:=0;j<3;j++{
				row:=(sqr/3)*3+i
				col:=(sqr%3)*3+j
				if board[row][col]=='.'{
					continue
				}
				if sqr_map[board[row][col]] {
					return false
				}
				sqr_map[board[row][col]] = true
			}
		}
	}
	return true
}
