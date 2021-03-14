package binary_trees_with_factors

import (
	"fmt"
	"sort"
)

func contains(s []int, e int) int {
	for i, a := range s {
		if a == e {
			return i
		}
	}
	return -1
}

func numFactoredBinaryTrees(arr []int) int {
	MOD := 1000000007
	sort.Ints(arr)
	fmt.Println(arr)
	dp := make([]int, len(arr))

	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[i] % arr[j] == 0 {
				fmt.Println(arr[i], arr[j] ,arr[i] / arr[j])
				find := contains(arr, arr[i] / arr[j])
				if find > -1 {
					dp[i] = (dp[i] + dp[j] * dp[find]) % MOD
				}
			}
		}
	}

	var result int
	for i := 0; i < len(dp); i++ {
		result = result + dp[i]
		fmt.Println(dp[i])
	}

	return result % MOD
}

func main() {
	fmt.Println(numFactoredBinaryTrees([]int{46,144,5040,4488,544,380,4410,34,11,5,3063808,5550,34496,12,540,28,18,13,2,1056,32710656,31,91872,23,26,240,18720,33,49,4,38,37,1457,3,799,557568,32,1400,47,10,20774,1296,9,21,92928,8704,29,2162,22,1883700,49588,1078,36,44,352,546,19,523370496,476,24,6000,42,30,8,16262400,61600,41,24150,1968,7056,7,35,16,87,20,2730,11616,10912,690,150,25,6,14,1689120,43,3128,27,197472,45,15,585,21645,39,40,2205,17,48,136}))
}

