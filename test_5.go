package main

import "fmt"

func main() {
	fmt.Println(sort([]int{1, 10, 2, 5, 9, 3, 4, 6, 8, 7}))
	fmt.Println(sort([]int{20, 42, 51, 2, 10, 11, 30}))
}

func sort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] < arr[j] {
				temp := arr[j]
				arr[j] = arr[i]
				arr[i] = temp
			}
		}
	}
	return arr
}
