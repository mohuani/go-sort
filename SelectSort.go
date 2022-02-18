package main

import "fmt"

// 选择排序：https://goa.lengthggirl.com/#/algorithm/sort/select_sort
// 逻辑思路：打扑克牌的时候，会习惯性地从左到右扫描，然后将最小的牌放在最左边，然后从第二张牌开始继续从左到右扫描第二小的牌，放在最小的牌右边，以此反复。选择排序和玩扑克时的排序特别相似。
// 编程思路：1、需要一个变量 min 存放一遍遍历后的最小值。2、需要一个变量 minIndex 存在这个这个最新小 min 在 一次遍历过程的位置  Index

func SelectSort(list []int) []int {
	length := len(list)

	for i := 0; i < length; i++ {
		min := list[i]
		minIndex := i

		for j := i + 1; j < length; j++ {
			if list[j] < min {
				min = list[j]
				minIndex = j
			}
		}

		// 最小值不是 i 的时候才进行交换
		if i != minIndex {
			list[i], list[minIndex] = list[minIndex], list[i]
		}
	}

	return list
}

func main() {
	list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	SelectSort(list)
	fmt.Println(list)
}
