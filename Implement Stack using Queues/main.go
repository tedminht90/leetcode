package main

import "fmt"

// Đây là định nghĩa một struct tên MyStack chứa 1 slice queue kiểu int
type MyStack struct {
	queue []int
}

// Đây là hàm khởi tạo trả về một MyStack mới
func Constructor() MyStack {
	return MyStack{}
}

// Đây là hàm thêm một phần tử vào cuối slice queue
func (this *MyStack) Push(x int)  {
	this.queue = append(this.queue, x)
}

func (this *MyStack) Pop() int {
	x := this.queue[len(this.queue)-1] // Lấy phần tử cuối cùng của slice queue
	this.queue = this.queue[:len(this.queue)-1] // Xóa phần tử cuối cùng của slice queue
	return x
}

// Chỉ trả về giá trị phần tử đỉnh stack mà không xoá nó
func (this *MyStack) Top() int {
	return this.queue[len(this.queue)-1]
}

// Kiểm tra xem stack có rỗng hay không
func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

func main() {
	input := []string{"MyStack","push","push","top","pop","empty"}

	results := make([]interface{},0)

	obj := Constructor()

	// Process each operation
    for i, op := range input {
        switch op {
        case "MyStack":
            results = append(results, nil) // thêm nil vào results
        case "push":
            if i == 1 { // Nếu i == 1 thì thêm 1 vào stack
                obj.Push(1)
                results = append(results, nil)
            } else if i == 2 { // Nếu i == 2 thì thêm 2 vào stack
                obj.Push(2)
                results = append(results, nil)
            }
        case "top":
            results = append(results, obj.Top()) // Lấy giá trị phần tử đỉnh stack
        case "pop":
            results = append(results, obj.Pop()) // Xoá phần tử đỉnh stack
        case "empty":
            results = append(results, obj.Empty()) // Kiểm tra stack có rỗng hay không
        }
    }
    fmt.Println(results)
}