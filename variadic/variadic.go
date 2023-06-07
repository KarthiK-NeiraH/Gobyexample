package main

import "fmt"

func main() {

	sum(1, 2)
	sum(1, 2, 3, 4, 5, 6, 7, 8)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}

func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0

	for index, num := range nums {
		fmt.Printf("index: %v\n", index)
		total += num
	}
	fmt.Println(total)
}
