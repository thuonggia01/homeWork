package main

import (
	"fmt"
	"math/rand"
	"time"
)

// BubbleSort
func bubbleSort(slice []int) []int {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	return slice
}

//InsertionSort
func insertionSort(slice []int) []int {
	var n = len(slice)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if slice[j-1] > slice[j] {
				slice[j-1], slice[j] = slice[j], slice[j-1]
			}
			j = j - 1
		}
	}
	return slice
}

//MergeSort
func mergeSort(items []int) []int {
	var num = len(items)

	if num == 1 {
		return items
	}

	middle := int(num / 2)
	var (
		left  = make([]int, middle)
		right = make([]int, num-middle)
	)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}

	return merge(mergeSort(left), mergeSort(right))
}
func merge(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}

//QuickSort
func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}

// Generates a slice
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(1)
	}
	return slice
}
func main() {
	slice := generateSlice(15)
	fmt.Println("slice before sort:\t	", slice)
	fmt.Println("slice after BubbleSort:  ", bubbleSort(slice))
	fmt.Println("****************************************")

	slice2 := generateSlice(15)
	fmt.Println("slice before sort:\t	", slice2)
	fmt.Println("slice after insertionSort:", insertionSort(slice2))
	fmt.Println("****************************************")
	slice3 := generateSlice(15)
	fmt.Println("slice before sort:\t	", slice3)
	fmt.Println("slice after MergeSort:", mergeSort(slice3))
	fmt.Println("****************************************")
	slice4 := generateSlice(15)
	fmt.Println("slice before sort:\t	", slice4)
	fmt.Println("slice after QuickSort:", quicksort(slice4))
	fmt.Println("****************************************")
}
