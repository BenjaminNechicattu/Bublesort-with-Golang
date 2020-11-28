package main

import (
	"fmt"
)

var toBeSorted [10]int = [10]int{1, 5, 2, 8, 8, 6, 9, 2, 7, 3}

func bubbleSort(input [10]int) {
	n := len(input)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if input[j] > input[i] {
				//fmt.Println("Swapping")
				input[i], input[j] = input[j], input[i]
			}
		}
	}
	fmt.Println(input)
}
func main() {
	fmt.Println("sorted array")
	bubbleSort(toBeSorted)
}
