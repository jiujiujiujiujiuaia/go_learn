package main

/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start

//没时间写代码了，讲讲思路

//目标：把有相同字母顺序的字符->转换成同一个key->记录在同一个坑中->返回

//解法一：遍历数组中所有的元素，然后排序该元素，然后丢到hash表中，那么相同的排序的元素
//都会放置在一个坑中，再遍历hash表就可以返回了
//时间复杂度n * klogk 空间复杂度n*k

//解法二：具有相同字符有两个条件，字符一样，字符数量一样。
//那么对于任意一个字符“abc”我们转换成1#1#1#0#0#....处理成这种情况即可
//时间复杂度为nk 空间复杂度为nk

func groupAnagrams(strs []string) [][]string {

}

// @lc code=end
