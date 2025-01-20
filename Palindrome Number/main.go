package main


func isPalindrome(x int) bool {
	n :=x 
	result := 0
	if x < 0 {
		return false
	} else {
		for x > 0{
			// lấy chữ số cuối cùng của x
			digit := x % 10
			// Thêm chữ số vào số đảo ngược
			result = result * 10 + digit
			// Loại bỏ chữ số cuối cùng của x
			x = x / 10
		}
	}
	if result != n {
		return false
	} else {
		return true
	}
}

func main(){
	x := 10
	println(isPalindrome(x))
}