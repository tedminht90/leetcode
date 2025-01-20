package main


func licenseKeyFormatting(S string, K int) string {
	n := len(S)
	var result string
	var count int
	for i := n-1; i >= 0; i-- {
		if S[i] == '-' {
			continue
		}
		if count == K {
			result = "-" + result
			count = 0
		}
		result = string(toUpper(S[i])) + result
		count++
	}
	return result
}

func toUpper(b byte) byte {
	if b >= 'a' && b <= 'z' {
		return b - 32
	}
	return b
}

func main(){
	S := "5F3Z-2e-9-w"
	K := 4
	println(licenseKeyFormatting(S, K))
}