/**
* @Author: Chao
* @Date: 2020-04-21 12:34
* @Version: 1.0
 */

package main

import "fmt"

func main() {
	nums := []int{2,2,2,1,2,2,1,2,2,2}
	//nums := []int{2, 4, 6}
	fmt.Println(numberOfSubarrays(nums, 2))
}

func numberOfSubarrays(nums []int, k int) int {
	n := len(nums)
	cnt := make([]int, n+1)
	odd := 0
	ans := 0
	cnt[0] = 1
	for i := 0; i < n; i++ {
		odd += nums[i]&1
		if odd >= k {
			ans += cnt[odd - k]
		}
		cnt[odd] += 1
		fmt.Println(odd, cnt[odd])
	}
	return ans
}