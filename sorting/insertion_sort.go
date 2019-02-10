package sorting

import (
	"log"
)

const (
	arrSize = 10
)

func InsertionSort(arr *[arrSize]int) {
	for i := 0; i <= arrSize-1; i++ {
		if i == 0 {
			continue
		}
		if arr[i] < arr[i-1] {
			temp := arr[i]
			arr[i] = arr[i-1]
			arr[i-1] = temp
		}

		for j := i; j >= len(arr[:j])-1; j-- {
			if j == 0 {
				break
			}
			if arr[j] < arr[j-1] {
				temp := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = temp
			} else {
				continue
			}

		}
	}
	log.Println("arr after insetion sort: ", arr)
}
