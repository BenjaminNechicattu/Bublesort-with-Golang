package main

import (
	"fmt"
)

//var toBeSorted []int = []int{1, 3, 2, 4, 8, 6, 5, 7, 2, 0}
var toBeSorted []int
var length int

func bubbleSort(input []int) {
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
	fmt.Println("enter array size : ")
	fmt.Scan(&length)
	fmt.Println(length)
	fmt.Println("enter array values ")
	toBeSorted := make([]int, length)
	for i := 0; i < length; i++ {
		//fmt.Println("enter next value :")
		fmt.Scan(&toBeSorted[i])
	}
	fmt.Println("sorted array :")
	bubbleSort(toBeSorted)
}
