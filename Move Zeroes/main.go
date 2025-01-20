package main

import "fmt"


func moveZeroes(nums []int) []int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	return nums
}

func main() {
	nums := []int{0,1,0,3,12}
	fmt.Println(moveZeroes(nums))
}

// Thuật toán sử dụng 2 con trỏ:

// i: Chỉ vị trí nơi số khác 0 tiếp theo sẽ được đặt
// j: Quét qua từng phần tử của mảng


// Cách hoạt động:

// Khi gặp số khác 0 (nums[j] != 0), ta hoán đổi phần tử tại vị trí i và j
// Sau khi hoán đổi, tăng i lên 1 để chuẩn bị cho vị trí tiếp theo
// j tiếp tục quét qua các phần tử còn lại