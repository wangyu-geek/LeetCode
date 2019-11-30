/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
// func climbStairs(n int) int {
// 	memSet := map[int]int{}
// 	return calcClimbStairs(0, n, memSet)
// }

//暴力递归
/*func calcClimbStairs(i, n int) int {
	if i > n {
		return 0
	}
	if i == n {
		return 1
	}
	return calcClimbStairs(i+1, n) + calcClimbStairs(i+2, n)
}*/

//记忆化递归
/*func calcClimbStairs(i, n int, memSet map[int]int) int {
	if i > n {
		return 0
	}
	if i == n {
		return 1
	}
	if memSet[i] > 0 {
		return memSet[i]
	}
	memSet[i] = calcClimbStairs(i+1, n, memSet) + calcClimbStairs(i+2, n, memSet)
	return memSet[i]
}*/

//动态规划
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	arr := []int{0, 1, 2}
	for i := 3; i <= n; i++ {
		arr = append(arr, arr[i-1] + arr[i-2])
	}
	return arr[n]
}


// @lc code=end