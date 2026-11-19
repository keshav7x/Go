package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, val := range nums {
		total += val
	}
	return total
}

func main() {
	// val := sum(1, 5, 8, 8, 10, 8, 8, 5)
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(slice...))
}
