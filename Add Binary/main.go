package main

import "fmt"

func addBinary(a,b string) string{
	var result string // Kết quả của chuỗi
	carry := 0 // biến nhớ
	i := len(a) - 1 // Vị trí cuối cùng của chuỗi a
	j := len(b) - 1 // Vị trí cuối cùng của chuỗi b

	for i >= 0 || j >= 0 {
		sum := carry
		if i >= 0 {
			// a[i] là một ký tự ('0' hoặc '1') trong chuỗi nhị phân
			// Trong ASCII:

			// '0' có giá trị là 48
			// '1' có giá trị là 49


			// Khi trừ '0':

			// '0' - '0' = 48 - 48 = 0
			// '1' - '0' = 49 - 48 = 1
			sum += int(a[i] - '0') //Chuyển ký tự '0' hoặc '1' thành số 0 hoặc 1
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0') //Chuyển ký tự '0' hoặc '1' thành số 0 hoặc 1
			fmt.Println(sum)
			j--
		}
		result = string(sum % 2 + '0') + result
		fmt.Println(result)
		carry = sum / 2
		fmt.Println(carry)
	}
	if carry > 0 {
		result = "1" + result
	}
	return result
}

func main() {
	a := "11"
	b := "1"
	fmt.Println(addBinary(a, b))
}

// a = "11"    (3 trong hệ thập phân)
// b = "1"     (1 trong hệ thập phân)

// Bước 1: i = 1, j = 0
// - sum = 0 + 1 + 1 = 2
// - result = "0", carry = 1

// Bước 2: i = 0, j = -1
// - sum = 1 + 1 = 2
// - result = "00", carry = 1

// Bước 3: i = -1, j = -1, carry = 1
// - Thêm carry vào đầu
// - Kết quả cuối: "100" (4 trong hệ thập phân)