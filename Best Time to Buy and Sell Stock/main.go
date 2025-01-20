package main

import "fmt"


func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	min := prices[0] // Giá thấp nhất đã thấy
	max := 0 // Lợi nhuận tối đa có thể đạt được
	for i := 1; i < len(prices); i++ { // Duyệt qua từng case để tìm ra lợi nhuận tối đa
		if prices[i] < min { // Nếu giá hiện tại nhỏ hơn giá thấp nhất đã thấy
			min = prices[i] // Cập nhật giá thấp nhất
		} else if prices[i] - min > max { // Nếu lợi nhuận hiện tại lớn hơn lợi nhuận tối đa đã thấy
			max = prices[i] - min // Cập nhật lợi nhuận tối đa
		}
	}
	return max
}

func main(){
	prices := []int{7,1,5,3,6,4}
	fmt.Println(maxProfit(prices))
}