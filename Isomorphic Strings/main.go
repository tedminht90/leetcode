package main


func isIsomorphic(s,t string) bool {
	if len(s) != len(t) {
		return false
	}

	m1 := make(map[byte]byte) // đảm bảo 1 ký tự trong s chỉ map với 1 ký tự trong t
	m2 := make(map[byte]byte) // đảm bảo 1 ký tự trong t chỉ map với 1 ký tự trong s

	for i := 0; i < len(s); i++ {
		// Kiểm tra mapping của s và t
		if _, ok := m1[s[i]]; !ok {
			m1[s[i]] = t[i] // Thêm mapping mới
		} else {
			if m1[s[i]] != t[i] { // Kiểm tra mapping cũ
				return false
			}
		}

		// Kiểm tra mapping của t và s
		if _, ok := m2[t[i]]; !ok {
			m2[t[i]] = s[i] // Thêm mappting mới
		} else {
			if m2[t[i]] != s[i] { // Kiểm tra mapping cũ
				return false
			}
		}
	}

	return true
}

func main(){
	s := "egg"
	t := "add"

	println(isIsomorphic(s, t))
}