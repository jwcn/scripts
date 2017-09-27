// 快速排序 Quick Sort in Golang
package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

func main() {

	// slice := generateSlice(20)
	slice := []int{3, 7, 2, 4, 5}
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	quicksort(slice)
	// foo(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")

}

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	// pivot := rand.Int() % len(a)

	// a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]
	fmt.Println("======", a[:left], a[left+1:])

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}

// test
func foo(a []int) []int {
	if len(a) < 2 {
		return a
	}

	base := a[0]
	// var left, right []int
	left := make([]int, 0)
	right := make([]int, 0)

	for i := 1; i < len(a); i++ {
		if a[i] < base {
			left = append(left, a[i])
		} else {
			right = append(right, a[i])
		}
	}

	l := foo(left)
	r := foo(right)

	b := []int{base}
	fmt.Println("==========", reflect.TypeOf(l), reflect.TypeOf(b), reflect.TypeOf(r))
	return l
	// return l + b + r // how to merge slice ? 
}
