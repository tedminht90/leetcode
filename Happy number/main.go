package main


func isHappy(n int) bool {
	if n == 1 {
		return true
	}
	m := make(map[int]bool) // lưu kết quả của các số đã xét
	for n != 1 {
		if m[n] { // nếu số đã xét trước đó thì return false
			return false
		}
		m[n] = true // lưu số đã xét
		n = sumOfSquares(n)
	}
	return true
}

func sumOfSquares(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}

func main(){
	n := 19
	println(isHappy(n))
}