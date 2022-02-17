package main

import "fmt"

// 冒泡排序：https://goa.lenggirl.com/#/algorithm/sort/bubble_sort

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
