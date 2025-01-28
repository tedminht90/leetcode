package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	// Kiểm tra 3 điều kiện
	// 1. flowerbed[i] == 0 -> chưa có hoa
	// 2. (i == 0 || flowerbed[i-1] == 0) -> hoa trước đó không có
	// 3. (i == len(flowerbed)-1 || flowerbed[i+1] == 0) -> vị trí sau trống hoặc là vị trí cuối cùng
	for i:=0; i<len(flowerbed); i++{
		if flowerbed[i] == 0 && 
		(i == 0 || flowerbed[i-1] == 0) && 
		(i == len(flowerbed)-1 || flowerbed[i+1] == 0){
			flowerbed[i] = 1
			count++
		}
	}
	return count >= n
}

func main(){
	flowerbed := []int{1,0,0,0,1}
	n := 1
	fmt.Println(canPlaceFlowers(flowerbed, n))
}