<?php
/*
 * @lc app=leetcode.cn id=1 lang=php
 *
 * [1] 两数之和
 */

// @lc code=start
class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $target
     * @return Integer[]
     */
    function twoSum($nums, $target) {
        $arrTarget = array();
        for ($i = 0; $i < count($nums); $i++) {
            $subTarget = $target - $nums[$i];
            //询问有没有人约我
            if (isset($arrTarget[$nums[$i]])) {
                return [ $arrTarget[$nums[$i]], $i ];
            } else {
                //没有人约,写下自己的要求
                $arrTarget[$subTarget] = $i;
            }
        }
        return [];
    }
}
// @lc code=end
