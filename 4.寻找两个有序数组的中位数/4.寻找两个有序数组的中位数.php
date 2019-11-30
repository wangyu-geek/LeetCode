<?php
/*
 * @lc app=leetcode.cn id=4 lang=php
 *
 * [4] 寻找两个有序数组的中位数
 */

// @lc code=start
class Solution {

    /**
     * @param Integer[] $nums1
     * @param Integer[] $nums2
     * @return Float
     */
    function findMedianSortedArrays($nums1, $nums2) {
        $nums1 = array_merge($nums1, $nums2);
        sort($nums1);

        $count = count($nums1);
        if ($count % 2 == 0) {
            $num_key_1 = floor($count / 2) - 1;
            $num_key_2 = floor($count / 2);
            $middle_num = ($nums1[$num_key_1] + $nums1[$num_key_2]) / 2;
        } else {
            $num_key = floor(($count - 1) / 2);
            $middle_num = $nums1[$num_key];
        }
        return $middle_num;
    }
}
// @lc code=end

