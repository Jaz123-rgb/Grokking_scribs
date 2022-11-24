package main

import "fmt"

// function responsible to perform binary search
func binarySearch(numbers []int, left int, right int, item int) (int, int) {
	if right >= left {
		mid := left + (right-left)/2
		if numbers[mid] == item {
			return numbers[mid], mid
		}
		if numbers[mid] > item {
			return binarySearch(numbers, left, mid-1, item)
		}
		return binarySearch(numbers, mid+1, right, item)
	}
	return -1, -1
}
func main() {
	numbers := []int{10, 20, 30, 40, 50}
	n := len(numbers)
	item := 40
	result, index := binarySearch(numbers, 0, n-1, item)
	if result == -1 && index == -1 {
		fmt.Println("Item is not present")
	} else {
		fmt.Println("Item", result, "is found at index", index)
	}
}
