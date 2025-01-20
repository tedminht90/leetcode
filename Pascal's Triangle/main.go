package main


import "fmt"

func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	if numRows == 0 {
		return result
	}
	result[0] = []int{1} // Hàng đầu tiên luôn có 1 phần tử là 1
	for i := 1; i < numRows; i++ {
		result[i] = make([]int, i + 1) // Mỗi hàng có i + 1 phần tử
		result[i][0] = 1 // Phần tử đầu tiên của mỗi hàng luôn là 1
		result[i][i] = 1 // Phần tử cuối cùng của mỗi hàng luôn là 1
		for j := 1; j < i; j++ {
			// Mỗi phần tử ở giữa bằng tổng của 2 phần tử ở hàng trên
			// result[i - 1][j - 1]: phần tử ở hàng trên, bên trái
			// result[i - 1][j]: phần tử ở hàng trên, bên phải
			result[i][j] = result[i - 1][j - 1] + result[i - 1][j]
		}
	}
	return result
}

func main() {
	numRows := 5
	fmt.Println(generate(numRows))
}