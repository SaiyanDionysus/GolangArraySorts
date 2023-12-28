package main

import (
	"fmt"
	"math/rand"
)

func mergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	firstSort := mergeSort(items[:len(items)/2])
	secondSort := mergeSort(items[len(items)/2:])
	return merge(firstSort, secondSort)
}

func merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}

func main() {
	unsortedArr := rand.Perm(10)
	fmt.Println("Unsorted random arr: ", unsortedArr)
	sorted := mergeSort(unsortedArr)
	fmt.Println("Sorted array: ", sorted)
}
