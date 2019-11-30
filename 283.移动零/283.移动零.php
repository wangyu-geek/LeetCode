/*
 * @lc app=leetcode.cn id=283 lang=php
 *
 * [283] 移动零
 */

// @lc code=start
class Solution {

    /**
     * @param Integer[] $nums
     * @return NULL
     */
    function moveZeroes(&$nums) {
        $nextNotZeroKey = 0; //保存数组中下一个值不为0的key下标
        $length = count($nums);
        for ($i = 0; $i < $length; $i++) {
            if ($nums[$i] != 0) {
                if ($nextNotZeroKey != $i) {
                    $nums[$nextNotZeroKey] = $nums[$i];
                    $nums[$i] = 0;
                }
                $nextNotZeroKey++;
            }
        }
    }
}
// @lc code=end

