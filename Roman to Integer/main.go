package main


func romanToInt(s string) int {
	romanMap := map[byte]int{
		'I':1,
		'V':5,
		'X':10,
		'L':50,
		'C':100,
		'D':500,
		'M':1000,
	}
	result := 0
	prevValue := 0
	// for i := 0; i < len(s); i++ {
	// 	if i+1 < len(s) && romanMap[s[i]] < romanMap[s[i+1]] {
	// 		result -= romanMap[s[i]]
	// 	}else{
	// 		result += romanMap[s[i]]
	// 	}
	// }
	
	for i := len(s) - 1; i >= 0; i-- {
        currentValue := romanMap[s[i]]
        if currentValue < prevValue {
            result -= currentValue
        } else {
            result += currentValue
        }
        prevValue = currentValue
    }
	return result
}

func main(){
	s := "VIVV"
	println(romanToInt(s))
}