package main

import "fmt"

// 归并排序：https://goa.lenggirl.com/#/algorithm/sort/merge_sort
//我们先介绍两个有序的数组合并成一个有序数组的操作。

func MergeSort2(array []int) []int {
	length := len(array)
	if length < 2 {
		return array
	}

	mid := length / 2
	leftArray := array[:mid]
	rightArray := array[mid:]

	return merge2(MergeSort2(leftArray), MergeSort2(rightArray))

}

// 归并操作
//1、先申请一个辅助数组，长度等于两个有序数组长度的和。
//2、从两个有序数组的第一位开始，比较两个元素，哪个数组的元素更小，那么该元素添加进辅助数组，然后该数组的元素变更为下一位，继续重复这个操作，直至数组没有元素。
//3、返回辅助数组。

func merge2(leftArray []int, rightArray []int) []int {
	var result []int

	// 左右数组都不为空的情况
	for len(leftArray) != 0 && len(rightArray) != 0 {
		if leftArray[0] <= rightArray[0] {
			result = append(result, leftArray[0])
			leftArray = leftArray[1:]
		} else {
			result = append(result, rightArray[0])
			rightArray = rightArray[1:]
		}
	}

	// 左数组有值，右数组无值的情况
	for len(leftArray) != 0 {
		result = append(result, leftArray[0])
		leftArray = leftArray[1:]
	}

	// 右数组有值，左数组无值的情况
	for len(rightArray) != 0 {
		result = append(result, rightArray[0])
		rightArray = rightArray[1:]
	}

	return result
}

// 自底向上的非递归实现

func main() {
	list := []int{5}
	fmt.Println(MergeSort2(list))

	list1 := []int{5, 9}
	fmt.Println(MergeSort2(list1))

	list2 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	fmt.Println(MergeSort2(list2))
}
