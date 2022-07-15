package main

import "fmt"

// 快速排序: 快速排序通过一趟排序将要排序的数据分割成独立的两部分，其中一部分的所有数据都比另外一部分的所有数据都要小，然后再按此方法对这两部分数据分别进行快速排序，整个排序过程可以递归进行，以此达到整个数据变成有序序列。

// 步骤如下：
// 先从数列中取出一个数作为基准数。一般取第一个数。
// 分区过程，将比这个数大的数全放到它的右边，小于或等于它的数全放到它的左边。
// 再对左右区间重复第二步，直到各区间只有一个数。

func main() {
	list := []int{5}
	fmt.Println(QuickSort3(list))

	list1 := []int{5, 9}
	fmt.Println(QuickSort3(list1))

	list2 := []int{5, 9, 1}
	fmt.Println(QuickSort3(list2))

	list3 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	fmt.Println(QuickSort3(list3))
}

//QuickSort3 参考php版本的快排写的一个golang版本
func QuickSort3(slice []int) []int {
	if cap(append(slice)) <= 1 {
		return slice
	}
	pivot := slice[0]

	var leftSlice []int
	var rightSlice []int

	for i := 1; i < len(slice); i++ {
		if slice[i] <= pivot {
			leftSlice = append(leftSlice, slice[i])
		} else {
			rightSlice = append(rightSlice, slice[i])
		}
	}

	leftSlice = QuickSort3(leftSlice)
	leftSlice = append(leftSlice, pivot)
	rightSlice = QuickSort3(rightSlice)

	return append(leftSlice, rightSlice...)
}
