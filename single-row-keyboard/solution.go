func calculateTime(keyboard string, word string) int {
	pointer := 0
	result := 0
	for i := 0; i < len(word); i++ {
		w := strings.Index(keyboard, string(word[i]))
		diff := w-pointer
		if diff > 0 {
			result = result + diff
		} else {
			result = result + diff*-1
		}
		pointer = w
	}
	return result
}
