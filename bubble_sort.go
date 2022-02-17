package main

// 冒泡排序

func BubbleSort(list []int) []int {
	len := len(list)

	for i := 0; i < len; i++ {
		for j := i +1 ; j < len; j++ {
			if list[i] > list[j] {
				tmp := list[i]
				list[i] = list[j]
				list[j] = tmp
			}
		}
	}

	return list
}







