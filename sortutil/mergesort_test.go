package sortutil

import (
	"testing"
	"math/rand"
	"sort"
)

var arr []struct {
	input, want []int
}

type MyInt []int

func (a MyInt) Len() int { return len(a) }
func (a MyInt) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a MyInt) Less(i, j int) bool { return a[i] < a[j] }

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

func init() {
	r := rand.New(rand.NewSource(99))
	count := r.Intn(99)
	arr = make([]struct {
		input, want []int
	}, count)
	for i := range arr {
		size := r.Intn(9999)
		arr[i].input = make([]int, size)
		for j := range arr[i].input {
			arr[i].input[j] = r.Intn(9999)
		}
		arr[i].want = arr[i].input
		sort.Sort(MyInt(arr[i].want))
	}
}