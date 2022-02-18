package main

import "fmt"

// 冒泡排序：https://goa.lenggirl.com/#/algorithm/sort/bubble_sort
// 逻辑思路：冒泡排序是大多数人学的第一种排序算法，在面试中，也是问的最多的一种，有时候还要求手写排序代码，因为比较简单。
//		   冒泡排序属于交换类的排序算法。
// 编程思路：

func BubbleSort(list []int) []int {
	length := len(list)

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if list[i] > list[j] {
				tmp := list[i]
				list[i] = list[j]
				list[j] = tmp
			}
		}
	}

	return list
}

func main() {
	list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	BubbleSort(list)
	fmt.Println(list)
}
