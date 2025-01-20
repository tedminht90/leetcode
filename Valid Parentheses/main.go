package main


func isValid(s string) bool {
	// Khởi tạo một slice stack rỗng
	stack := []rune{}
	// Map ánh xạ từ đầu đóng sang dấu mở tương ứng
	m := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	// Duyệt qua từng ký tự trong chuỗi
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			// Nếu là ký tự mở ngoặc thì thêm vào stack
			stack = append(stack, c)
		} else if len(stack) > 0 && stack[len(stack)-1] == m[c] {
			// Kiếm tra 2 điều kiện:
			// 1. stack không rỗng (len(stack) > 0)
			// 2. dấu mở ngoặc trên cùng của stack (stack[len(stack)-1])
			//   khớp với dấu mở ngoặc tương ứng của ký tự đóng ngoặc (m[c])
			// thì loại bỏ ký tự trên cùng của stack
			stack = stack[:len(stack)-1]
		} else {
			// Trường hợp còn lại: không hợp lệ
			return false
		}
	}
	// Nếu stack rỗng thì chuỗi hợp lệ

	return len(stack) == 0

}

func main() {
	s := "({[]})"
	//s := "({[)}]"
	//s := ")]}([{)"
	println(isValid(s))
}

/*
Chuỗi: "({[]})"

Bước 1: '('  -> stack = ['(']
Bước 2: '{'  -> stack = ['(', '{']
Bước 3: '['  -> stack = ['(', '{', '[']
Bước 4: ']'  -> khớp với '[' -> stack = ['(', '{']
Bước 5: '}'  -> khớp với '{' -> stack = ['(']
Bước 6: ')'  -> khớp với '(' -> stack = []
Kết quả: true (stack rỗng)

Chuỗi: "([)]"

Bước 1: '('  -> stack = ['(']
Bước 2: '['  -> stack = ['(', '[']
Bước 3: ')'  -> không khớp với '[' -> return false
*/