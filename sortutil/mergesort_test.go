package sortutil

import "testing"

var arr = []struct {
	input, want []int
}{
	{
		input: []int{6, 202, 100, 301, 38, 8, 1},
		want:  []int{1, 6, 8, 38, 100, 202, 301},
	},
	{
		input: []int{1},
		want:  []int{1},
	},
	{
		input: []int{},
		want:  []int{},
	},
}

// TestMergeSortRecursive
func TestMergeSortRecursive(t *testing.T) {
	for _, obj := range arr {
		got := MergeSortRecursive(obj.input)
		if !compare(got, obj.want) {
			t.Errorf("MergeSortRecursive(%v) == %v, want %v\n", obj.input, got, obj.want)
		}
	}
}

// TestMergeSort
func TestMergeSort(t *testing.T) {
	for _, obj := range arr {
		got := MergeSort(obj.input)
		if !compare(got, obj.want) {
			t.Errorf("MergeSort(%v) == %v, want %v\n", obj.input, got, obj.want)
		}
	}
}

// 比较两个 []int 是否相等
func compare(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
