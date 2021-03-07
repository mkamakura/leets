func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums) - 1; i++ {
		for j := i + 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]
			if sum == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
