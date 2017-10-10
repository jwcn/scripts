// 冒泡排序 Bubble Sort in Golang
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// slice := generateSlice(20)
	slice := []int{6, 4, 2, 5, 7}
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	bubbleSort(slice)
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

/**
又叫插入排序?
1. 从第一位开始与右侧比, 大就互换位置
2. 当一轮循环下来没有换过一次位置, 就结束最外围的死循环
*/
func bubbleSort(items []int) {
	var (
		n      = len(items)
		sorted = false
	)
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if items[i] > items[i+1] { // 与右侧比, 大就互换位置
				items[i+1], items[i] = items[i], items[i+1]
				swapped = true
			}
		}

		if !swapped {
			sorted = true // 一轮循环下来没有换过一次位置, 就结束最外围的死循环
		}
		n = n - 1 // 排除上一轮循环出来的最右最大值后, 开始下一轮循环
	}
}

/*
又称选择排序(Selection Sort)
1. 从右侧第一位开始与左侧每一位对比,小就换
2. 第二次循环对比排除最左侧第一位, 第三次排除第二位, 依次类推
*/
func foo(items []int) []int {
	length := len(items)
	if length < 2 {
		return items
	}

	for i := 0; i < length; i++ {
		for j := length - 1; j > i; j-- {
			if items[j] < items[j-1] {
				items[j], items[j-1] = items[j-1], items[j]
			}
		}
	}

	return items
}
