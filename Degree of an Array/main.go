package main

import "fmt"

// left: vị trí đầu tiên xuất hiện
// right: vị trí cuối cùng xuất hiện
// count: số lần xuất hiện
type Info struct {
    left, right, count int
}

func findShortestSubArray(nums []int) int {
    info := make(map[int]*Info) // Map lưu thông tin của mỗi số
    maxDegree := 0 // Biến lưu degree lớn nhất

    // Duyệt một lần để lấy tất cả thông tin
    for i, num := range nums {
        if _, exists := info[num]; !exists {
			// Số xuất hiện lần đầu
            info[num] = &Info{left: i, right: i, count: 1}
        } else {
			// Số đã xuất hiện trước đó
            inf := info[num]
            inf.right = i // Cập nhật vị trí cuối cùng
            inf.count++ // Tăng số lần xuất hiện
        }
        maxDegree = max(maxDegree, info[num].count)
    }

    // Tìm độ dài ngắn nhất
    result := len(nums) // Khởi tại kết quả = độ dài mảng
    for _, inf := range info {
        if inf.count == maxDegree {
			// Nếu số có count = maxDegree
			// Tính độ dài đoạn chứa tất cả xuất hiện của số đó
            result = min(result, inf.right - inf.left + 1)
        }
    }

    return result
}

func main(){
	nums := []int{1,2,2,3,1}
	fmt.Println(findShortestSubArray(nums))
}