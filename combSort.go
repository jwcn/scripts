package main

import "fmt"

func main() {
	slice := []int{6, 4, 2, 5, 7}
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	CombSort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}

func CombSort(items []int) {
	var (
		shrink  = 1.3
		swapped = true

		counter = len(items)
		gap     = len(items)
	)

	for swapped {
		swapped = false

		gap = int(float64(gap) / shrink)
		for i := 0; i+gap < counter; i++ {
			if items[i] > items[i+gap] {
				items[i], items[i+gap] = items[i+gap], items[i]
				swapped = true
			}
		}

	}
}
