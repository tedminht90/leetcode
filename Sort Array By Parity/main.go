package main


func sortArrayByParity(nums []int) []int{
	l, r := 0, len(nums)-1  // Con trỏ trái và phải
	for l < r {  // Chạy khi l chưa vượt quá r
		if nums[l] % 2 == 1 {  // Nếu phần tử bên trái lẻ
			nums[l], nums[r] = nums[r], nums[l]  // Đổi chỗ với phần tử bên phải
			r--  // Giảm con trỏ bên phải
		} else {
			l++  // Tăng con trỏ bên trái
		}
	}
	return nums
}


func main() {
	nums := []int{3, 1, 2, 4}
	println(sortArrayByParity(nums))
}