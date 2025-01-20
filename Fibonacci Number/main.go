package main

func fib(n int) int {
	if n <= 1 {
		return n
	}
	
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a + b
	}
	
	return b
}


func main(){
	num := 4
	println(fib(num))
}