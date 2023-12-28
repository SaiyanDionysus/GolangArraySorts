package main

import (
	"fmt"
	"math/rand"
)

func BubbleSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

func main() {
	array := rand.Perm(10)
	fmt.Println("Your unsorted array: ", array)
	fmt.Println("Your sorted array:", BubbleSort(array))
}
