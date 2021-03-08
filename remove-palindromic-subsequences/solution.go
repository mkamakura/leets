package remove_palindromic_subsequences

func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i,j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

func removePalindromeSub(s string) int {
	if len(s) == 0 {
		return 0
	}

	if s == reverse(s) {
		return 1
	}

	return 2
}
