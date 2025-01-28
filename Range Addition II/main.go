package main

func maxCount(m int, n int, ops [][]int) int {
	// Nếu không có phép tăng nào
	if len(ops) == 0 {
		return m*n // Toàn ma trận là 0
	}

	// Tìm vị trí nhỏ nhất củ các phép tăng
	minA, minB := ops[0][0], ops[0][1]
	for i := 1; i < len(ops); i++ {
		if ops[i][0] < minA {
			minA = ops[i][0]
		}
		if ops[i][1] < minB {
			minB = ops[i][1]
		}
	}
	// số phần tử lớn nhất =  vùng được tăng nhiều nhất
	return minA * minB
}

func main() {
	m:= 3
	n:= 3
	ops:= [][]int{{2,2},{3,3}}
	println(maxCount(m,n,ops))	
}