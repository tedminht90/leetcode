package main


func addStrings(num1 string, num2 string) string {
	i := len(num1) - 1
	j := len(num2) - 1
	carry := 0 // biến nhớ
	res := "" // kết quả

	for i >= 0 || j >= 0 || carry != 0 { 
		if i >= 0 {
			carry += int(num1[i] - '0') // chuyển kí tự số về số
			i-- 
		}
		if j >= 0 {
			carry += int(num2[j] - '0')
			j--
		}
		res = string('0' + carry%10) + res
		carry /= 10
	}

	return res
}

func main(){
	num1 := "11"
	num2 := "99"

	println(addStrings(num1, num2))
}