package main

/*
 * @lc app=leetcode.cn id=85 lang=golang
 *
 * [85] 最大矩形
 */

// @lc code=start

//解法一，暴力法，假设二维数组长n宽m
//那么找到任意两点作为对角点：(nm)^2
//然后确定这个范围内没有非1的点然后计算面积
//(nm)^2 * (mn)

//解法二：优化暴力法 （mn^2）
//把这个题目转换成84 也就说记录宽度，然后横着堪称一个矩形问题

//解法三：记录好宽度，然后把这道题目转换成84 n复杂度

func maximalRectangle(matrix [][]byte) int {

}

// @lc code=end
