package main

import "fmt"

// 快速排序: 快速排序通过一趟排序将要排序的数据分割成独立的两部分，其中一部分的所有数据都比另外一部分的所有数据都要小，然后再按此方法对这两部分数据分别进行快速排序，整个排序过程可以递归进行，以此达到整个数据变成有序序列。

// 步骤如下：
// 先从数列中取出一个数作为基准数。一般取第一个数。
// 分区过程，将比这个数大的数全放到它的右边，小于或等于它的数全放到它的左边。
// 再对左右区间重复第二步，直到各区间只有一个数。

// 参考文章：https://selfboot.cn/2016/09/01/lost_partition/

func main() {
	list := []int{5}
	QuickSort(list, 0, len(list)-1)
	fmt.Println(list)

	list1 := []int{5, 9}
	QuickSort(list1, 0, len(list1)-1)
	fmt.Println(list1)

	list2 := []int{5, 9, 1}
	QuickSort(list2, 0, len(list2)-1)
	fmt.Println(list2)

	list3 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	QuickSort(list3, 0, len(list3)-1)
	fmt.Println(list3)
}

func QuickSort(array []int, begin, end int) {
	if begin < end {
		loc := partition(array, begin, end)

		QuickSort(array, begin, loc-1) // 左方排序
		QuickSort(array, loc+1, end)   // 右方排序
	}
}

// 快速排序中用到的 partition 算法思想很简单，首先从无序数组中选出枢轴点 pivot，然后通过一趟扫描，
//以 pivot 为分界线将数组中其他元素分为两部分，使得左边部分的数小于等于枢轴，右边部分的数大于等于枢轴（左部分或者右部分都可能为空），最后返回枢轴在新的数组中的位置。
func partition(array []int, begin, end int) int {
	i := begin + 1
	j := end

	// 已基准数base将数组分成两部分，小于base的和大于base的
	for i < j {
		if array[i] > array[j] {
			array[i], array[j] = array[j], array[i]
			j--
		} else {
			i++
		}
	}

	// 此时的数组被分成  (array[begin+1] ~ array[i]) < array[begin] < (array[j] ~ array[end])
	if array[i] >= array[begin] {
		i--
	}

	array[begin], array[i] = array[i], array[begin]

	// 这里为什么要return i，i的含义是  枢轴在新的数组中的位置
	return i
}
