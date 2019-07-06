package main

import "fmt"

func find(x int, slice []int) int {
	i, j := 0, len(slice)-1

	for i <= j {
		mid := int((i + j) / 2)
		if x == slice[mid] {
			return slice[mid]
		}
		if x < slice[mid] {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}

	return 0
}

func main() {
	var slice []int
	for i := 1; i <= 1000000; i++ {
		slice = append(slice, i)
	}

	fmt.Println(find(8270, slice))
	fmt.Println(find(1, slice))
	fmt.Println(find(999999, slice))
}
