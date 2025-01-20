package main

func checkPerfectNumber(num int) bool {
    if num <= 0 {
        return false
    }
    
    sum := 0
    // Duyệt từ 1 đến num/2 vì không có ước số nào lớn hơn num/2 
    for i := 1; i <= num/2; i++ {
        if num % i == 0 { // i là ước số của num
            sum += i // cộng ước số vào sum
        }
    }
    
    return sum == num
}


func main(){
	num := 28
	println(checkPerfectNumber(num))
}