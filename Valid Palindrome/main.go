package main


func isPalindrome(s string) bool {
	l, r := 0, len(s)-1 // Khởi tạo 2 con trỏ l, r (đầu và cuối chuỗi)
	for l < r { // chạy khi con trỏ bên trái nhỏ hơn con trỏ bên phải
		if !isAlphanumeric(s[l]) {
			l++ // Bỏ qua, di chuyển con trỏ trái
		} else if !isAlphanumeric(s[r]) {
			r-- // Bỏ qua, di chuyển con trỏ phải
		} else if toLower(s[l]) != toLower(s[r]) {
			return false // Nếu 2 ký tự không giống nhau thì trả về false
		} else {
			l++ // Di chuyển cả hai con trỏ
			r--
		}
	}
	return true
}

func isAlphanumeric(c byte) bool {
	return 	(c >= '0' && c <= '9') || 	// Là số
			(c >= 'a' && c <= 'z') || 	// Là chữ thường
			(c >= 'A' && c <= 'Z') 		// Là chữ hoa
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32 // Chuyển chữ hoa thành chữ thường
	}
	return c
}

func main(){
	s := "A man, a plan, a canal: Panama"
	println(isPalindrome(s))
}