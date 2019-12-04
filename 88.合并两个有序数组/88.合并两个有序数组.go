package main

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
type SortIntSlice []int

func (s SortIntSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s SortIntSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortIntSlice) Len() int {
	return len(s)
}

/*
 * 方法一,合并数组然后排序
 */
/*
func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:m], nums2...)
	sort.Sort(SortIntSlice(nums1))
}*/

/*
 * 方法二，双指针
 */
func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1_temp := make([]int, m, m)
	copy(nums1_temp, nums1)
	pointer1, pointer2, i := 0, 0, 0
	for (pointer1 < m) && (pointer2 < n) {
		if nums1_temp[pointer1] < nums2[pointer2] {
			nums1[i] = nums1_temp[pointer1]
			pointer1++
		} else {
			nums1[i] = nums2[pointer2]
			pointer2++
		}
		i++
	}
	for pointer1 < m {
		nums1[i] = nums1_temp[pointer1]
		i++
		pointer1++
	}

	for pointer2 < n {
		nums1[i] = nums2[pointer2]
		i++
		pointer2++
	}

}

// @lc code=end

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}
