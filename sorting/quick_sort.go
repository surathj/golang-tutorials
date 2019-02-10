package sorting

import (
	"log"
)

func QuickSort(arr *[arrSize]int, low, high int) {
	if low < high {
		pi := Partition(arr, low, high)
		QuickSort(arr, low, pi-1)
		QuickSort(arr, pi+1, high)
	}
	//log.Println("arr after quick sort:", arr)
}

func Partition(arr *[arrSize]int, low, high int) int {
	i := low - 1
	pivot := arr[high]

	for j := low; j <= high-1; j++ {
		// if left counter is less than pivot
		if arr[j] <= pivot {
			//increase wall
			i++
			//swap wall and j
			temp := arr[j]
			arr[j] = arr[i]
			arr[i] = temp
		}
	}
	temp2 := arr[high]
	arr[high] = arr[i+1]
	arr[i+1] = temp2
	log.Println("array: ", arr)
	return i + 1
}
