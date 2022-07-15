package main

import "fmt"

// 归并排序：https://goa.lenggirl.com/#/algorithm/sort/merge_sort
//我们先介绍两个有序的数组合并成一个有序数组的操作。

// 一、 自顶向下归并排序，排序范围在 [begin,end) 的数组

func MergeSort(array []int, begin int, end int) {
	if end-begin > 1 {

		// 将数组一分为二，分为 array[begin,mid) 和 array[mid,high)
		mid := begin + (end-begin+1)/2

		// 左边数组排序
		MergeSort(array, begin, mid)

		// 右边数组排序
		MergeSort(array, mid, end)

		// 归并两个已经排序好的数组
		merge(array, begin, mid, end)
	}

}

// 归并操作
//1、先申请一个辅助数组，长度等于两个有序数组长度的和。
//2、从两个有序数组的第一位开始，比较两个元素，哪个数组的元素更小，那么该元素添加进辅助数组，然后该数组的元素变更为下一位，继续重复这个操作，直至数组没有元素。
//3、返回辅助数组。
func merge(array []int, begin int, mid int, end int) []int {

	leftSize := mid - begin
	rightSize := end - mid
	newSize := leftSize + rightSize // 辅助数组的长度
	result := make([]int, 0, newSize)

	leftStep, rightStep := 0, 0 // 数组切割的步长

	for leftStep < leftSize && rightStep < rightSize {
		leftValue := array[begin+leftStep]
		rightValue := array[mid+rightStep]

		if leftValue < rightValue {
			result = append(result, leftValue)
			leftStep++
		} else {
			result = append(result, rightValue)
			rightStep++
		}
	}

	// 将剩下的数组追加到辅助数组里面
	result = append(result, array[begin+leftStep:mid]...)
	result = append(result, array[mid+rightStep:end]...)

	// 将辅助数组的元素复制回原数组，这样该辅助空间就可以被释放掉
	for i := 0; i < newSize; i++ {
		array[i] = result[i]
	}

	return array
}

// 自底向上的非递归实现

func main() {
	list := []int{5}
	MergeSort(list, 0, len(list))
	fmt.Println(list)

	list1 := []int{5, 9}
	MergeSort(list1, 0, len(list1))
	fmt.Println(list1)

	list2 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	MergeSort(list2, 0, len(list2))
	fmt.Println(list2)
}
