package main


func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	// Kiểm tra xem có giao nhau trên trục x và y
    xOverlap := !(rec1[2] <= rec2[0] || rec1[0] >= rec2[2])
    yOverlap := !(rec1[3] <= rec2[1] || rec1[1] >= rec2[3])
	return xOverlap && yOverlap
}

func main(){
	rec1 := []int{0,0,2,2}
	rec2 := []int{1,1,3,3}
	println(isRectangleOverlap(rec1, rec2))
}