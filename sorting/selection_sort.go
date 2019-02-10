package sorting

import (
	"log"
)

func SelectionSort(arr *[arrSize]int, low int) {
	//{7, 2, 1, 8, 6, 3, 5, 4}
	//{6, 1, 0, 7, 5, 2, 4, 3}
	//{2,1,5,7,6}
	//	jIndex := 0
	for i := 0; i <= arrSize-1; i++ {
		slice := arr[i:]
		minIndex := FindMin(slice, i)
		temp := arr[i]
		arr[i] = arr[minIndex]
		arr[minIndex] = temp

		//log.Println("min: ", arr[i])
	}
	log.Println("arr after selection sort: ", arr)
}

func FindMin(arr []int, originalIndex int) int {
	min := 1000000
	minIndex := 0
	for i := 0; i <= len(arr)-1; i++ {
		if arr[i] < min {
			min = arr[i]
			minIndex = i
		}
	}
	return minIndex + originalIndex
}
