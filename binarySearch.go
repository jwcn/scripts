// 二分查找 Binary Search in Golang
package main

import "fmt"

func binarySearch(needle int, haystack []int) int {

	low := 0
	high := len(haystack) - 1

	for low <= high {
		median := (low + high) / 2

		if haystack[median] < needle {
			low = median + 1
		} else if haystack[median] > needle {
			high = median - 1
		} else {
			return median
		}
	}

	return -1
}

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(binarySearch(63, items))
}
