/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	numsMap := map[int]int{}
	for key,value := range nums {
		sub := target - value
		if anotherValue, ok := numsMap[sub]; ok {
			return []int{anotherValue, key}
		}
		numsMap[value] = key
	}
	return []int{}
}
// @lc code=end

//func main(){
////	res := twoSum([]int{3, 3}, 6)
////	fmt.Println(res)
////}
