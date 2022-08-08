package binarySearch

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"normal case：mid 任意时刻不是上半区有序就是下半区有序，直到找到或者不断缩小半区" +
			"直到缩小到一个元素", args{[]int{5, 6, 7, 1, 2, 3, 4}, 1}, 3},
		{"only two items: 测试nums[l]<`=` nums[mid]", args{[]int{1, 0}, 1}, 0},
		{"only one items: 测试找不到循环退出的情况", args{[]int{0}, 1}, -1},
		{"degrade to binary search: ", args{[]int{1, 2, 3, 4, 5, 6, 7}, 1}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
