/**
* @Author: Chao
* @Date: 2020/4/23 13:24
* @Version: 1.0
 */

package main

import "fmt"

func main() {
	fmt.Println(waysToChange(25))
}

func waysToChange(n int) int {
	dp := make([]int, n + 1)
	coins := []int{1, 5, 10, 25}
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= n; i++ {
			dp[i] = (dp[i] + dp[i - coin]) % 1000000007
		}
		//fmt.Println(dp)
	}
	return dp[n]
}