import (
	"sort"
	"math"
)
/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个有序数组的中位数
 */

// @lc code=start
type IntSlice []int

func (p IntSlice) Len() int {
	return len(p)
}

func (p IntSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p IntSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

//N为奇数时，m0.5 = X((N+1)/2) 偶数时，m0.5 = (X(N/2) + X(N/2+1)) / 2
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	length := len(nums1)
	sort.Sort( IntSlice(nums1) )

	if length % 2 == 0 {
		num1_key := int( math.Floor(  float64((length - 1) / 2)) )
		num2_key := int(math.Floor(float64(length / 2)))
		return float64( (nums1[num1_key] + nums1[num2_key])) / 2
	} else {
		num_key := int(math.Floor(float64((length - 1)/2) ))
		return float64(nums1[num_key])
	}
}
// @lc code=end