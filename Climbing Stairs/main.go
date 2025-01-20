package main

func climbStairs(n int) int {
	if n <=2 {
		return n
	}

	one := 1 // số cách leo 1 bậc
	two := 2 // số cách leo 2 bậc

	for i := 3; i <= n; i++ {
		temp := one + two
		one = two
		two = temp
	}
	return two
}


func main() {
	n := 5
	println(climbStairs(n))
}