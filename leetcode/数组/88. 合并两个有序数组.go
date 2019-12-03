package main

func main() {
	merge([]int{2, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	tailIndex := len(nums1) - 1
	index1, index2 := m-1, n-1
	for ; tailIndex > -1; tailIndex-- {
		if index1 == -1 && index2 != -1 {
			nums1[tailIndex] = nums2[index2]
			index2--
		} else if index2 == -1 && index1 != -1 {
			index1--
		} else if nums1[index1] > nums2[index2] {
			swap(nums1[:], tailIndex, index1)
			index1--
		} else {
			nums1[tailIndex] = nums2[index2]
			index2--
		}
	}
}

func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}
