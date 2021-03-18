func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func wiggleMaxLength(nums []int) int {
    if len(nums) < 2 {
        return len(nums)
    }

    up := make([]int, len(nums))
    down := make([]int, len(nums))

    up[0] = 1
    down[0] = 1

    for i := 1; i < len(nums); i++ {
        if nums[i] > nums[i-1] {
            up[i] = down[i-1] + 1
            down[i] = down[i-1]
        } else if nums[i] < nums[i-1] {
            down[i] = up[i-1] + 1
            up[i] = up[i-1]
        } else {
            down[i] = down[i-1]
            up[i] = up[i-1]
        }
    }

    return max(down[len(nums) - 1], up[len(nums) - 1])
}
