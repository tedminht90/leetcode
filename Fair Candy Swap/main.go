package main

import "fmt"


func fairCandySwap(aliceSizes, bobSizes []int) []int {
    // Tính tổng ban đầu
    sumA, sumB := 0, 0
    for _, v := range aliceSizes {
        sumA += v
    }
    for _, v := range bobSizes {
        sumB += v
    }

    // Lưu các giá trị của Bob vào map
    m := make(map[int]int)  // Đổi thành map[int]int để đếm số lần xuất hiện
    for _, v := range bobSizes {
        m[v]++
    }

    // Với mỗi hộp của Alice, tìm hộp tương ứng của Bob
    for _, x := range aliceSizes {
        // Giải phương trình: (sumA - x + y) = (sumB - y + x)
        // → 2y = 2x + (sumB - sumA)
        // → y = x + (sumB - sumA)/2
        y := x + (sumB - sumA)/2
        if m[y] > 0 {
            // Kiểm tra xem sau khi trao đổi có cân bằng không
            newSumA := sumA - x + y
            newSumB := sumB - y + x
            if newSumA == newSumB {
                return []int{x, y}
            }
        }
    }

    return []int{}  // Không tìm thấy cặp số thỏa mãn
}

func main(){
	aliceSizes := []int{1, 1}
	bobSizes := []int{2, 3}
	fmt.Println(fairCandySwap(aliceSizes, bobSizes))
}