package main

import "fmt"

// 插入排序：https://goa.lenggirl.com/#/algorithm/sort/insert_sort
// 逻辑思路：每次把一个数插到已经排好序的数列里面形成新的排好序的数列，以此反复。
//         有些人打扑克时习惯从第二张牌开始，和第一张牌比较，第二张牌如果比第一张牌小那么插入到第一张牌前面，这样前两张牌都排好序了，
//         接着从第三张牌开始，将它插入到已排好序的前两张牌里，形成三张排好序的牌，后面第四张牌继续插入到前面已排好序的三张牌里，直至排序完。
// 编程思路：

func InsertSort(list []int) []int {

	length := len(list)
	for i := 1; i < length-1; i++ {
		deal := list[i] // 需要排序的数字，下边从 i=1 开始
		j := i - 1      // i 左边的数字

		if deal < list[j] {
			for ; j >= 0 && deal < list[j]; j-- {
				list[j+1] = list[j]
			}

			list[j+1] = deal
		}

	}

	return list
}

func main() {
	list := []int{5}
	InsertSort(list)
	fmt.Println(list)

	list1 := []int{5, 9}
	InsertSort(list1)
	fmt.Println(list1)

	list2 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	InsertSort(list2)
	fmt.Println(list2)
}
