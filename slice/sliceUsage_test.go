package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sliceCapAndLen(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}

	length, capacity := sliceLenAndCap(slice)
	assert.Equal(t, 5, length)
	assert.Equal(t, 5, capacity)

	initLength, initCapacity := 5, 10
	slice = make([]int, initLength, initCapacity)
	length, capacity = sliceLenAndCap(slice)
	assert.Equal(t, initLength, length)
	assert.Equal(t, initCapacity, capacity)
}

func Test_sliceCapAndLen_makeSlice(t *testing.T) {
	t.Log("if don't give the concrete capacity, the default capacity equal length")
	initLength := 5
	slice := make([]int, initLength)
	length, capacity := sliceLenAndCap(slice)
	assert.Equal(t, initLength, length)
	assert.Equal(t, initLength, capacity)
}

func Test_sliceCapAndLen_shareSlice(t *testing.T) {
	t.Log("if slice2 shares directly slice1, the slice2 length = end - start, the slice2 capacity = cap(slice1) - start")
	slice := []int{1, 2, 3, 4, 5}
	start := 1
	end := 3
	slice1 := slice[start:end]
	length, capacity := sliceLenAndCap(slice1)
	assert.Equal(t, end-start, length)
	assert.Equal(t, cap(slice)-start, capacity)
}

func Test_sliceUsedInFunction(t *testing.T) {
	slice1 := []int{1, 2, 3, 4, 5}
	newValue := 100

	t.Logf("before append operation, slice1=%v", slice1)
	sliceUsedInFunction(slice1, newValue)
	t.Logf("before append operation, slice1=%v", slice1)
}

func Test_sliceUsedInFunction_EnoughCapacity(t *testing.T) {
	t.Log("在底层数组capacity充足的情况下，slice1和slice2 共享同一个底层数组，所以底层数组地址是一样的，但是视图不一样")
	t.Log("在函数间传递中，是值传递，因此切片的地址不一样，但是切片的值是一样的（底层数组）")
	t.Log("尽管slice1和slice2 内容不同，但是仍然共享一个底层数组")
	slice1 := make([]int, 1, 2)
	newValue := 100

	t.Logf("before append operation, slice1=%v bottom array addr=%p", slice1, &slice1[0])
	slice2 := sliceAppendOperation(slice1, newValue)
	t.Logf("after append operation, slice1=%v bottom array addr=%p", slice1, &slice1[0])
	t.Logf("after append operation, slice2=%v bottom array addr=%p", slice2, &slice2[0])

	assert.Equal(t, len(slice1), 1)
	assert.NotEqual(t, slice1, slice2)

}

func Test_sliceUsedInFunction_NotEnoughCapacity(t *testing.T) {
	t.Log("在底层数组capacity不充足的情况下，slice1和slice2不再共享同一个切片，所以底层数组发生了改变")
	t.Log("在函数间传递中，是值传递，因此切片的地址不一样，但是切片的值是一样的（底层数组）")
	t.Log("slice1和slice2 内容不同，也不共享一个底层数组")
	slice1 := make([]int, 1, 1)
	newValue := 100

	t.Logf("before append operation, slice1=%v bottom array addr=%p", slice1, &slice1[0])
	slice2 := sliceAppendOperation(slice1, newValue)
	t.Logf("after append operation, slice1=%v bottom array addr=%p", slice1, &slice1[0])
	t.Logf("after append operation, slice2=%v bottom array addr=%p", slice2, &slice2[0])

	assert.Equal(t, len(slice1), 1)
}

func Test_understandAppendOperation_notEnoughCapacity(t *testing.T) {
	t.Log("append的操作是直接在slice末尾加上一个元素，如果slice的capacity不够，则创建一个新的底层数组")
	base := make([]int, 0, 1)
	base = append(base, 100)
	slice1 := append(base, 200)
	slice2 := append(base, 300)

	assert.Equal(t, 200, slice1[1])
	assert.Equal(t, 300, slice2[1])
	assert.NotEqual(t, &base[0], &slice1[0])
	assert.NotEqual(t, &base[0], &slice2[0])
	t.Logf("addr base=%p,slice1=%p,slice2=%p", &base[0], slice1, slice2)
}

func Test_understandAppendOperation_enoughCapacity(t *testing.T) {
	t.Log("append的操作是直接在slice末尾加上一个元素，如果slice的capacity够，那就是在原有的底层数组增加")
	t.Log("1.Warning: 如果多个slice都在一个base slice进行了append，由于没有创建新的底层数组，会出现覆盖的情况")
	base := make([]int, 0, 2)
	base = append(base, 100)
	slice1 := append(base, 200)
	slice2 := append(base, 300)

	assert.Equal(t, 300, slice1[1])
	assert.Equal(t, 300, slice2[1])
	assert.Equal(t, &base[0], &slice1[0])
	assert.Equal(t, &base[0], &slice2[0])
	t.Logf("addr base=%p,slice1=%p,slice2=%p", &base[0], slice1, slice2)
}

func Test_sliceShardArrayImpactEachOther_success(t *testing.T) {
	t.Log("slice1 and slice2 share the same bottom array, if slice2 modify Nth item and not exceed cap, in fact, " +
		"modify the Nth item in the bottom array")
	modifiedIndex, modifiedValue := 1, 100
	start, end := 1, 3
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := slice1[start:end]

	t.Logf("before modify operation, slice1=%v,slice2=%v", slice1, slice2)
	sliceShardArrayImpactEachOther(slice2, modifiedIndex, modifiedValue)
	t.Logf("after modify operation, slice1=%v,slice2=%v", slice1, slice2)
	assert.Equal(t, modifiedValue, slice2[modifiedIndex])
	assert.Equal(t, modifiedValue, slice1[modifiedIndex+start])

}

func Test_sliceShardArrayImpactEachOtherByAppending_Success(t *testing.T) {
	t.Log("slice1 and slice2 share the same bottom array, if slice2 append new item and not exceed cap, in fact, " +
		"modify the Nth item in the bottom array")
	newValue := 100
	start, end := 1, 3
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := slice1[start:end]

	t.Logf("before append operation, slice1=%v,slice2=%v", slice1, slice2)
	slice2 = sliceShardArrayImpactEachOtherByAppending(slice2, newValue)
	t.Logf("after append operation, slice1=%v,slice2=%v", slice1, slice2)
	assert.Equal(t, newValue, slice2[end-start])
	assert.Equal(t, newValue, slice1[end])
}

func Test_sliceShardArrayImpactEachOtherByAppending_declareSliceCapacity(t *testing.T) {
	t.Log("approach: keep new slice cap and length same")
	t.Log("declare/limit capacity for new slice, if exceed the capacity, will create new array for new slice")
	newValue := 100
	start, end := 1, 3
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := slice1[start:end:3]

	t.Logf("before append operation, slice1=%v,slice2=%v \n", slice1, slice2)
	slice2 = sliceShardArrayImpactEachOtherByAppending(slice2, newValue)
	t.Logf("after append operation, slice1=%v,slice2=%v \n", slice1, slice2)
	assert.Equal(t, newValue, slice2[end-start])
	assert.NotEqual(t, newValue, slice1[end])
}

func Test_sliceShardArrayImpactEachOtherByAppending_declareZeroSlice(t *testing.T) {
	t.Log("create new slice with zero capacity by making operation not share, new slice has its own bottom array, then copy from dst" +
		"to src slice")
	newValue := 100
	start, end := 1, 3
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, 0)
	slice2 = append(slice2, slice1[start:end]...)

	t.Logf("before append operation, slice1=%v,slice2=%v", slice1, slice2)
	slice2 = sliceShardArrayImpactEachOtherByAppending(slice2, newValue)
	t.Logf("after append operation, slice1=%v,slice2=%v", slice1, slice2)
	assert.Equal(t, newValue, slice2[end-start])
	assert.NotEqual(t, newValue, slice1[end])
}

func Test_sliceShardArrayImpactEachOtherByAppending_declareNewSliceByMaking(t *testing.T) {
	t.Log("create new slice by making operation not share, new slice has its own bottom array, then copy from dst" +
		"to src slice")
	newValue := 100
	start, end := 1, 3
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, end-start)
	copy(slice2, slice1[1:3])

	t.Logf("before append operation, slice1=%v,slice2=%v", slice1, slice2)
	slice2 = sliceShardArrayImpactEachOtherByAppending(slice2, newValue)
	t.Logf("after append operation, slice1=%v,slice2=%v", slice1, slice2)
	assert.Equal(t, newValue, slice2[end-start])
	assert.NotEqual(t, newValue, slice1[end])
}
