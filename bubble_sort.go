package main

import (
	"fmt"
)

var toBeSorted [10]int = [10]int{1, 3, 2, 4, 8, 6, 5, 7, 2, 0}

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
