package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	// Khởi tạo các biến cần thiết
	n := len(prices) // Số lượng ngày giao dịch
	k := 2 // Số lượng giao dịch tối đa
	dp := make([][]int, k+1) // Mảng lưu trữ lợi nhuận tối đa

	// dp[0] = make([]int, n) // Mảng cho 0 giao dịch
	// dp[1] = make([]int, n) // Mảng cho tối đa 1 giao dịch
	// dp[2] = make([]int, n) // Mảng cho tối đa 2 giao dịch
	for i := 0; i <= k; i++ { // i là số lượng giao dịch
		dp[i] = make([]int, n)
	}
	// Duyệt qua từng case để tìm ra lợi nhuận tối đa
	// Với i=1: tính lợi nhuận tối đa khi giao dịch 1 lần
	// Với i=2: tính lợi nhuận tối đa khi giao dịch 2 lần
	for i := 1; i <= k; i++ {
		// profit đại diện cho lợi nhuận tốt nhất nếu mua cổ phiếu
		// Khởi tạo bằng -prices[0] giả sử mua vào ngày đầu tiên
		profit := -prices[0]
		for j := 1; j < n; j++ { //j là ngày giao dịch, duyệt từ 1 đến n-1
			// Tại ngày j, có 2 lựa chọn
			// 1. Giữ nguyên lợi nhuận từ trước: dp[i][j-1]
			// 2. Bán cổ phiếu vào ngày j: profit+prices[j]
			dp[i][j] = max(dp[i][j-1], profit+prices[j])
			// Có mua cổ phiếu vào ngày j không
			// Không mua: giữ progit cũ
			// Mua mới: lấy lợi nhuận của (i-1) giao dịch trừ đi giá mua hôm nay
			profit = max(profit, dp[i-1][j]-prices[j])
		}
	}
	return dp[k][n-1]
}

func main(){
	prices := []int{3,3,5,0,0,3,1,4}
	fmt.Println(maxProfit(prices))
}