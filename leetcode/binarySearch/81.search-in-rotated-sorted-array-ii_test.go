package binarySearch

import "testing"

func Test_search4(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"重复元素case", args{nums: []int{1, 0, 1, 1, 1}, target: 0}, true},
		{"重复元素case", args{nums: []int{1, 1, 1, 0, 1}, target: 0}, true},
		{"空数组", args{nums: []int{}, target: 3}, false},
		{"原地旋转", args{nums: []int{1, 2, 3, 4, 5}, target: 6}, false},
		{"原地旋转", args{nums: []int{1, 2, 3, 4, 5}, target: 5}, true},
		{"退化成O(N)", args{nums: []int{1, 1, 1, 1, 1}, target: 2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search4(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search4() = %v, want %v", got, tt.want)
			}
		})
	}
}
