package main

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	slow := 0 //慢指针
	for quick := 1; quick < len(nums); quick++ {
		if nums[slow] != nums[quick] {
			slow++
			nums[slow] = nums[quick]
		}
	}
	return slow + 1
}

// @lc code=end

func main() {
	arr := []int{1, 1, 2}
	res := removeDuplicates(arr)
	fmt.Println(arr)
	fmt.Println(res)
}
