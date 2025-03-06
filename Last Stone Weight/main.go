package main

import (
	"sort"
)


func lastStoneWeight(stones []int) int{
	for len(stones) > 1 {
		// Sắp xếp mảng giảm dần
		sort.Slice(stones, func(i, j int) bool {
			return stones[i] > stones[j]
		})

		// Lấy 2 viên đá nặng nhất
		x := stones[0]
		y := stones[1]
		// Nếu 2 viên đá bằng nhau, xóa cả 2 viên đá
		if x == y {
			stones = stones[2:]
		} else {
			// Nếu 2 viên đá khác nhau, thêm vào mảng viên đá mới với trọng lượng là hiệu 2 viên đá
			stones = stones[2:]
			stones = append(stones, x-y)
		}
	}
	// Nếu mảng rỗng trả về 0, ngược lại trả về phần tử còn lại
	if len(stones) == 0 {
		return 0
	}
	return stones[0]
}

func main(){
	stones := []int{2,7,4,1,8,1}
	println(lastStoneWeight(stones))
}