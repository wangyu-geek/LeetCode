/*
 * @lc app=leetcode.cn id=11 lang=php
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
class Solution {

    /**
     * @param Integer[] $height
     * @return Integer
     */
    function maxArea($height) {
        $i = 0;
        $j = count($height) - 1;
        $area = 0;

        while ($i < $j) {
            $area = max($area, min($height[$i], $height[$j]) * ($j - $i) );
            if ($height[$i] > $height[$j]) {
                $j--;
            } else {
                $i++;
            }
        }
        return $area;
    }
}
// @lc code=end

