package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start

//暴力求解1
//func threeSum(nums []int) [][]int {
//	team := [][]int{}
//	for i := 0; i < len(nums) - 2; i++ { //每个人
//		for j := i + 1; j < len(nums) - 1; j++ { //依次拉上其他人
//			for k := j + 1; k < len(nums); k++ { //问剩下的人
//				if nums[i] + nums[j] + nums[k] == 0 { //能不能组队
//					team = append(team, []int{nums[i], nums[j], nums[k]})
//				}
//			}
//		}
//	}
//	return team
//}

//暴力求解2
//func threeSum(nums []int) [][]int {
//	team := [][]int{}
//	registerMap := map[int][]int{}
//
//	for i := 0; i < len(nums) - 1; i++ { //每个人
//		for j := i + 1; j < len(nums); j++ { //依次拉上其他人
//			if registerTeam, ok := registerMap[nums[j]]; ok { //有没有适合自己的队伍
//				registerTeam = append([]int{nums[j]}, registerTeam...)
//				team = append(team, registerTeam)
//				delete(registerMap, nums[j]) // 从登记表删除
//			} else {
//				targetNum := -(nums[i] + nums[j])
//				registerMap[targetNum] = []int{nums[i], nums[j]} //写入登记表
//			}
//		}
//	}
//	return team
//}

type SortIntSlice []int

func (intSlice SortIntSlice) Len() int {
	return len(intSlice)
}

func (intSlice SortIntSlice) Less(i, j int) bool {
	return intSlice[i] < intSlice[j]
}

func (intSlice SortIntSlice) Swap(i, j int) {
	intSlice[i], intSlice[j] = intSlice[j], intSlice[i]
}

//先排序，后求解
func threeSum(nums []int) [][]int {
	term := [][]int{}
	if len(nums) < 3 {
		return term
	}
	//排序
	sort.Sort(SortIntSlice(nums))

	for i := range nums[:len(nums)-2] {
		if nums[i] > 0 {
			return term
		}
		//去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				term = append(term, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				//去重
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if sum < 0 {
				//值较小，左指针向右移
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			} else {
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			}
		}
	}
	return term
}

// @lc code=end

func main() {
	arr := []int{2, -3, 0, -2, -5, -5, -4, 1, 2, -2, 2, 0, 2, -4}
	res := threeSum(arr)
	fmt.Println(res)
}
