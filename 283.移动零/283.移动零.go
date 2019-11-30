package _83_移动零

import "fmt"

/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int) {
	notZeroKey := 0
	if len(nums) <= 1 {
		return
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if notZeroKey != i {
				nums[notZeroKey] = nums[i]
				nums[i] = 0
			}
			notZeroKey++
		}
	}
}

// @lc code=end

func main() {
	nums := []int{0, 1}
	moveZeroes(nums)
	fmt.Println(nums)
}
