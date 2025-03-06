package main

func validBoomerang(points [][]int) bool {

	x1, y1 := points[0][0], points[0][1]
	x2, y2 := points[1][0], points[1][1]
	x3, y3 := points[2][0], points[2][1]

	// Kiểm tra 3 điểm có trùng nhau không
	if x1 == x2 && y1 == y2 {
		return false
	}
	if x1 == x3 && y1 == y3 {
		return false
	}
	if x2 == x3 && y2 == y3 {
		return false
	}

	// Kiểm tra 3 điểm có cùng nằm trên 1 đường thẳng đứng hoặc ngang hay không
	if (x1 == x2 && x2 == x3) || (y1 == y2 && y2 == y3) {
		return false
	}

	// Kiểm tra 3 điểm thẳng hàng dùng độ dốc
	// ct (y₂-y₁)*(x₃-x₁) == (y₃-y₁)*(x₂-x₁)
	if (y2-y1)*(x3-x1) == (y3-y1)*(x2-x1) {
		return false
	}
	return true
}

func  main(){
	points := [][]int{{1, 1}, {2, 3}, {3, 2}}
	println(validBoomerang(points))
}