package main

import "fmt"


func countPrimeSetBits(left, right int) int{
	cnt := 0
	// Duyệt qua các số từ left đến right
	for i := left; i <= right; i++ {
		// Đếm số bit 1 và kiểm tra xem có phải số nguyên tố không
		if isPrime(bitsCount(i)) {
			cnt++
		}
	}
	return cnt
}

// Hàm đếm bit 1 (bitsCount)
func bitsCount(num int) int {
	cnt := 0
	for num > 0 {
		cnt += num & 1 // Cộng bit cuối vào kết quả
		num >>= 1 // Dịch phải 1 bit
	}
	return cnt
}

func isPrime(num int) bool {
	// Xử lý các trường hợp đặc biệt
	if num == 1 {
		return false
	}
	if num == 2 || num == 3 {
		return true
	}
	if num % 2 == 0 || num % 3 == 0 {
		return false
	}

	// Kiểm tra từ 5 đến căn bậc 2 của num
	// Tối ưu bằng cách tăng i lên 6 đơn vị mỗi lần
	for i := 5; i * i <= num; i += 6 {
		if num % i == 0 || num % (i + 2) == 0 {
			return false
		}
	}
	return true
}

func main(){
	left := 6
	right := 10

	fmt.Println(countPrimeSetBits(left, right))
}