package check_if_a_string_contains_all_binary_codes_of_size_k

func pow(base, k int) int {
	result := 1
	for i := 0; i < k; i++ {
		result = result * 2
	}
	return result
}

func hasAllCodes(s string, k int) bool {
	if k > len(s) {
		return false
	}

	count := pow(2, k)
	m := make(map[string]bool)
	for i := 0; i < len(s) - k + 1; i++ {
		m[s[i:i + k]] = true
		if len(m) == count {
			return true
		}
	}
	return false
}

