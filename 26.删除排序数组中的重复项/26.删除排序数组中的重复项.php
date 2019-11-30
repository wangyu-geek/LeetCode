<?php
/*
 * @lc app=leetcode.cn id=26 lang=php
 *
 * [26] 删除排序数组中的重复项
 */

// @lc code=start
class Solution {

    //方法二
    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function removeDuplicates(&$nums) {
        $slow = 0;
        for ($quick = 1; $quick < count($nums); $quick++) {
            if ($nums[$slow] != $nums[$quick]) {
                $slow++;
                $nums[$slow] = $nums[$quick];
            }
        }
        return $slow+1;
    }
}
// @lc code=end

$arr = [1,1,2];
$obj = new Solution();
$obj->removeDuplicates($arr);
var_dump($arr);