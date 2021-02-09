package main

import "fmt"

//https://leetcode-cn.com/problems/valid-sudoku/
func unitValid(i1, i2, j1, j2 int, board [][]byte)bool{
	validMap := make(map[byte]int)
	for i := i1; i <= i2; i++{
		for j := j1; j <= j2; j++{
			if board[i][j] >= '0' && board[i][j] <= '9'{
				if _, ok := validMap[board[i][j]]; ok{
					return false
				}
				validMap[board[i][j]] = 1
			}
		}
	}
	return true
}

func isValidSudoku(board [][]byte) bool {

	for i := 0; i < 9; i++{
		if !unitValid(i, i, 0, 8, board) ||
			!unitValid(0, 8, i, i, board){
			return false
		}
	}

	for i := 0; i < 3; i++{
		for j := 0; j < 3; j++{
			valid := unitValid(3*i, 3*i+2, 3*j, 3*j+2, board)
			if !valid{
				return false
			}
		}
	}
	return true
}

func main(){
	board :=[][]byte{
	{'.','.','4','.','.','.','6','3','.'},
	{'.','.','.','.','.','.','.','.','.'},
	{'5','.','.','.','.','.','.','9','.'},
	{'.','.','.','5','6','.','.','.','.'},
	{'4','.','3','.','.','.','.','.','1'},
	{'.','.','.','7','.','.','.','.','.'},
	{'.','.','.','5','.','.','.','.','.'},
	{'.','.','.','.','.','.','.','.','.'},
	{'.','.','.','.','.','.','.','.','.'},
	}

	fmt.Printf("isvalid :%v\n", isValidSudoku(board))
}
