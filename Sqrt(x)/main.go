package main


func mySqrt(x int) int {
    if x <= 1 { 
		return x 
	}
    left, right := 1, x
    for left < right {
        mid := left + (right-left)/2 // tương đương với mid = (left+right)/2, viết kiểu này để tránh tràn số khi left và right lớn
        if mid <= x/mid { // nếu mid^2 <= x, viết kiểu này để tránh tràn số khi mid^2 lớn
            left = mid + 1
        } else {
            right = mid
        }
    }
    return right - 1
}

func main(){
	x := 8
	println(mySqrt(x))
}
