package main

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=435 lang=golang
 *
 * [435] 无重叠区间
 */

// @lc code=start
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	return len(intervals) - merge(intervals)
}

type Interval struct {
	Start int
	End   int
}

type Intervals []Interval

func (Interval Intervals) Len() int {
	return len(Interval)
}

func (Interval Intervals) Swap(i, j int) {
	Interval[i], Interval[j] = Interval[j], Interval[i]
}

func (Interval Intervals) Less(i, j int) bool {
	return Interval[i].Start < Interval[j].Start
}

//算出来了有多少个不相交区间
func merge(intervals [][]int) int {

	//排序
	intervals = tranformArrayToStruct(intervals)

	minEnd := intervals[0][1]
	total := 1
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= minEnd {
			minEnd = max(minEnd, intervals[i][1])
			total++
		} else {
			minEnd = min(minEnd, intervals[i][1])
		}
	}
	return total
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//就是转个换排个序
func tranformArrayToStruct(intervals [][]int) [][]int {
	var Inter Intervals

	for i := 0; i < len(intervals); i++ {
		Inter = append(Inter, Interval{Start: intervals[i][0], End: intervals[i][1]})
	}

	sort.Sort(Inter)

	res := make([][]int, 0)
	for _, inter := range Inter {
		res = append(res, []int{inter.Start, inter.End})
	}

	return res
}

// @lc code=end
