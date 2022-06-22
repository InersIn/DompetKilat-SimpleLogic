package main

func reverse(s string) string {
	bstr := []rune(s)
	for x := 0; x < len(bstr)/2; x++ {
		bstr[len(bstr)-x-1], bstr[x] = bstr[x], bstr[len(bstr)-x-1]
	}
	return string(bstr)
}

func IsPalindrome(str string) bool {
	var status bool = true
	nStr := reverse(str)

	for i, v := range str {
		if byte(v) != nStr[i] {
			status = false
			break
		}
	}
	return status
}
