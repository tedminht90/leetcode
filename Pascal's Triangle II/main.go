package main

import "fmt"


func getRow(rowIndex int) []int {
	result := make([]int, rowIndex + 1)
	result[0] = 1
	for i := 1; i <= rowIndex; i++ {
		result[i] = result[i - 1] * (rowIndex - i + 1) / i
	}
	return result
}


func main(){
	rowIndex := 4
	fmt.Println(getRow(rowIndex))
}