package main

import "fmt"

func findRestaurant(list1, list2 []string) []string {
	m := make(map[string]int)
	for i, v := range list1 {
		m[v] = i // lưu vị trí của các phần tử trong list1
	}
	minSum := len(list1) + len(list2) // giả sử tổng lớn nhất có thể
	var res []string // mảng kết quả
	for i, v := range list2 {
		if j, ok := m[v]; ok { // kiểm tra nhà hàng có trong list1 không, nếu có
			if i+j < minSum { // nếu tổng vị trí nhỏ hơn tổng nhỏ nhất đã tìm thấy
				minSum = i + j // cập nhật tổng nhỏ nhất
				res = []string{v} // cập nhật mảng kết quả
			} else if i+j == minSum { // Cùng tổng nhỏ nhất
				res = append(res, v) // Thêm vào kết quả
			}
		}
	}
	return res
}


func main() {
	list1:= []string{"Shogun","KFC"}
	list2:= []string{"KFC","Shogun"}
	// list1:= []string{"Shogun","Tapioca Express","Burger King","KFC"}
	// list2:= []string{"Piatti","The Grill at Torrey Pines","Hungry Hunter Steakhouse","Shogun"}
	fmt.Println(findRestaurant(list1,list2))	
}