package strobogrammatic_number

func isStrobogrammatic(num string) bool {
	left := 0
	right := len(num)-1

	for left <= right {
		if num[left] == num[right] {
			if num[left] == 49 || num[left] == 56 || num[left] == 48 {
				left++
				right--
			} else {
				return false
			}
		} else if num[left] == 54 && num[right] == 57 || num[left] == 57 && num[right] == 54 {
			left++
			right--
		} else {
			return false
		}
	}

	return true
}
