package main

import "fmt"


func merge(num1, num2 []int, n,m int ){
	// Khởi tạo con trỏ cho num1, num2, và kết quả được merged
	i := m - 1 // Con trỏ cuối cùng của num1
	j := n - 1 // Con trỏ cuối cùng của num2
	k := m + n - 1 // Con trỏ cuối cùng của kết quả merged

	// Lặp qua từng phần tử của num1 và num2
	for i >= 0 && j >= 0 {
		if num1[i] > num2[j] {
			num1[k] = num1[i]
			i--
		} else {
			num1[k] = num2[j]
			j--
		}
		k--
	}

	// Nếu num2 còn phần tử thì copy vào num1
	for j >= 0 {
		num1[k] = num2[j]
		j--
		k--
	}

}


func main() {
	nums1 := []int{1,2,3,0,0,0}
	m := 3
	nums2 := []int{2,5,6}
	n := 3
	merge(nums1, nums2, n,m)
	fmt.Printf("%v", nums1)
}

// Bước 1: nums1[2]=3 vs nums2[2]=6
// - 3 < 6
// - nums1[5] = 6
// - j--, k--
// [1,2,3,0,0,6]

// Bước 2: nums1[2]=3 vs nums2[1]=5
// - 3 < 5
// - nums1[4] = 5
// - j--, k--
// [1,2,3,0,5,6]

// Bước 3: nums1[2]=3 vs nums2[0]=2
// - 3 > 2
// - nums1[3] = 3
// - i--, k--
// [1,2,3,3,5,6]

// Bước 4: nums1[1]=2 vs nums2[0]=2
// - 2 = 2 
// - nums1[2] = 2
// - j--, k--
// [1,2,2,3,5,6]

// Bước 5: nums1[1]=2 vs nums2[0]=2
// - 2 = 2
// - nums1[1] = 2 
// - i--, k--
// [1,2,2,3,5,6]

// Bước 6: Chỉ còn nums1[0]=1
// - nums1[0] = 1
// [1,2,2,3,5,6]