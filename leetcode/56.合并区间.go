package main

import "sort"

/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start

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

func merge(intervals [][]int) [][]int {
	res := make([][]int, 0)
	
	//排序
	intervals = tranformArrayToStruct(intervals)

	for i := 0; i < len(intervals); i++ {
		//前一个end小于后一个start 说明不在一个区间
		if len(res) == 0 || res[len(res)-1][1] < intervals[i][0] {
			res = append(res, intervals[i])
		} else {
			res[len(res)-1][1] = max(intervals[i][1], res[len(res)-1][1])
		}
	}
	return res
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
