// 归并排序：将已有序的子序列合并，得到完全有序的序列；
//     速度仅次于快速排序，为稳定排序算法。
//
// 原理：
//     0. 初始序列大小为 size = 2
//     1. 将数组以 size 为大小分割，形成 ceil(n/size) 个序列，对每个序列进行归并操作
//     2. size = size * 2
//     3. 重复步骤 1-2，直到 size >= 2n
//
// 归并操作原理：
//     0. 申请空间，大小为两个已经排序序列之和，用来存放合并后的序列
//        result := make([]int, 0, len(left)+len(right))
//     1. 两个指针，分别指向两个已排序序列的起始位置
//     2. 比较两个指针所指的元素，把相对较小的元素，放到 result 中，并且指针移动到下一位
//     3. 重复步骤 2 ，直到某一指针超过序列尾
//     4. 把另一序列剩余元素，合并到 result 尾部
//
// 例子：
//     假设序列 {6, 202, 100, 301, 38, 8, 1}
//     1. 第一次归并（比较 3 次）; size = 2, 序列数 = 4
//        {6, 202}, {100, 301}, {8, 38}, {1}
//     2. 第二次归并（比较 4 次）; size = 4, 序列数 = 2
//        {6, 100, 202, 301}, {1, 8, 38}
//     3. 第三次归并（比较 4 次）; size = 8, 序列数 = 1
//        {1, 6, 8, 38, 100, 202, 301}
package sortutil

// MergeSortRecursive 通过递归进行归并排序
func MergeSortRecursive(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	middle := length / 2
	left := MergeSortRecursive(arr[:middle])
	right := MergeSortRecursive(arr[middle:])
	return merge(left, right)
}

// MergeSort 通过循环进行归并排序
func MergeSort(arr []int) []int {
	size := 2
	arrLen := len(arr)
	for ; size <= arrLen*2; size *= 2 {
		part := arrLen / size
		if arrLen%size != 0 {
			part += 1
		}
		length := size / 2
		var newArr []int
		for i := 0; i < part; i++ {
			begin := i * size
			end := (i + 1) * size
			if end > arrLen {
				end = arrLen
			}
			subarr := arr[begin:end]
			left := subarr[:length]
			right := subarr[length:]
			newArr = append(newArr, merge(left, right)...)
		}
		arr = newArr
	}
	return arr
}

// merge 归并操作，合并两个数组
func merge(left []int, right []int) []int {
	l, r := 0, 0
	result := make([]int, 0, len(left)+len(right))
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return result
}
