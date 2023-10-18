package solutions

func IsValidSudoku(board [][]byte) bool {
	var rowMap [9]map[byte]bool
	var colMap [9]map[byte]bool
	var blockMap [9]map[byte]bool

	for i := 0; i < 9; i++ {
		rowMap[i] = map[byte]bool{}
		colMap[i] = map[byte]bool{}
		blockMap[i] = map[byte]bool{}
	}

	var num byte
	var blockIndex int
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			num = board[row][col]
			if num == '.' {
				continue
			}

			blockIndex = row/3*3 + col/3
			if rowMap[col][num] || colMap[row][num] || blockMap[blockIndex][num] {
				return false
			}

			rowMap[col][num] = true
			colMap[row][num] = true
			blockMap[blockIndex][num] = true
		}
	}

	return true
}
