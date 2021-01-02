package utils

func ReverseString(str string) string {
	n := 0
	runeArr := make([]rune, len(str))
	for _, r := range str {
		runeArr[n] = r
		n++
	}

	for i := 0; i < n/2; i++ {
		runeArr[i], runeArr[n-1-i] = runeArr[n-1-i], runeArr[i]
	}

	return string(runeArr)
}
