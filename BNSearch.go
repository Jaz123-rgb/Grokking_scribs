package main

import "fmt"

func search(nums []int, target int) int {
	var i, j = 0, len(nums) - 1
	for i <= j {
		var mid = i + (j-i)/2
		if nums[mid] == target {
			return mid
		}
		if target > nums[mid] {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return -1
}
func main() {
	numbers := []int{10, 20, 30, 40, 50}
	item := 40
	fmt.Println(search(numbers, item))

}
