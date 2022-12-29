package slice

func sliceLenAndCap(slice []int) (int, int) {
	return len(slice), cap(slice)
}

func sliceShardArrayImpactEachOther(slice2 []int, modifiedIndex, modifiedValue int) {
	slice2[modifiedIndex] = modifiedValue
}

func sliceShardArrayImpactEachOtherByAppending(slice []int, newValue int) []int {
	slice = append(slice, newValue)
	return slice
}

func sliceUsedInFunction(slice []int, newValue int) {
	slice[0] = newValue
}

func sliceAppendOperation(slice []int, newValue int) []int {
	slice = append(slice, newValue)
	return slice
}
