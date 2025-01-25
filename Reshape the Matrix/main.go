package main

import "fmt"

// len(mat[0]) là số cột của ma trận
func matrixReshape(mat [][]int, r int, c int) [][]int {
	// kiểm tra điều kiện tổng số phàn tử giữ nguyên
	if len(mat) * len(mat[0]) != r *c {
		return mat
	}
	// Khởi tạo ma trận mới với kích thước r x c
	res := make([][]int, r)
	for i := range res {
		res[i] = make([]int, c)
	}
	// Duyệt từng phần tử matrix gốc
	for i:=0; i<len(mat); i++ {
		for j:=0; j<len(mat[0]); j++ {
			// Tính toán vị trí 1 chiều của phần tử trong ma trận mới
			pos := i*len(mat[0]) + j
			// Tính toán vị trí hàng và cột của phần tử trong ma trận mới
			newRow := pos / c
			newCol := pos % c
			res[newRow][newCol] = mat[i][j]
		}
	}
	return res
}


func main(){
	mat := [][]int{{1,2},{3,4}}
	r := 2
	c := 4
	fmt.Println(matrixReshape(mat, r, c))
}