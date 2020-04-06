package main

import "container/heap"

/*
 * @lc app=leetcode.cn id=295 lang=golang
 *
 * [295] 数据流的中位数
 */

// @lc code=start

//解法一：维护一个不排序的数组，插入的时候o(1)，寻找的时候可以用215寻找第k大的数。那么就是o(n)

//解法二：维护一个排序的数组，插入的时候平均是0(n),寻找的时候o(1)

//解法三：二叉搜索树，插入的时候记录一下属于第几个点，插入的时候是logn，寻找的时候也是logn
//但是如果不平衡的话可以退化成链表

//解法四：维护两个堆，一个最大堆一个最小堆，最大堆是数组前半部分，最小堆是数组后半部分。
//我们要达到的效果是总数为偶数，最大堆（l）和最小堆（h）数量相等，总数为奇数，最大堆比最小堆多一个
//步骤：先加入到l中，然后把l中最大的给h，如果size h比l大了，再把h的最小给l
//数量比较l h
//       1 0
//       1 1
//       2 1
//       2 2
//通过上面你就可以看到这两个堆的数量是平衡的

//golang中的堆要自己搞
type IntMinHeap []int

func (h IntMinHeap) Len() int           { return len(h) }
func (h IntMinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntMinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h IntMinHeap) Top() int {
	return h[0]
}

type IntMaxHeap []int

func (h IntMaxHeap) Len() int           { return len(h) }
func (h IntMaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntMaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h IntMaxHeap) Top() int {
	return h[0]
}

type MedianFinder struct {
	MaxHeap *IntMaxHeap
	LowHeap *IntMinHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	finder := MedianFinder{MaxHeap: new(IntMaxHeap), LowHeap: new(IntMinHeap)}
	heap.Init(finder.MaxHeap)
	heap.Init(finder.LowHeap)
	return finder
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.MaxHeap, num)
	top := heap.Pop(this.MaxHeap).(int)
	heap.Push(this.LowHeap, top)

	if this.LowHeap.Len() > this.MaxHeap.Len() {
		top = heap.Pop(this.LowHeap).(int)
		heap.Push(this.MaxHeap, top)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if (this.MaxHeap.Len()+this.LowHeap.Len())%2 == 0 {
		return (float64(this.MaxHeap.Top()) + float64(this.LowHeap.Top())) / 2
	} else {
		return float64(this.MaxHeap.Top())
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
// @lc code=end
