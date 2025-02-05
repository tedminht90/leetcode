package main

import "fmt"


func judgeCircle(moves string) bool {
	x := 0  // Toạ độ trục x
	y := 0 // Toạ độ trục y

	// Duyệt từng bước di chuyển
	for _, move := range moves {
		switch move {
		case 'U': y++ // Lên trên
		case 'D': y-- // Xuống dưới
		case 'L': x-- // Sang trái
		case 'R': x++ // Sang phải
		}
	}

	// Kiểm tra có về vị trí ban đầu không
	return x == 0 && y == 0
}


func main(){
	fmt.Println(judgeCircle("UD"))
	fmt.Println(judgeCircle("LL"))
}