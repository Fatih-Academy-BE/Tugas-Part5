package main

import "fmt"

func RemoveData(arrayNumber []int) []int {
	// your code here
}

func main() {
	fmt.Println(RemoveData([]int{2, 3, 3, 3, 4, 5}))       // [2 3 4 5]
	fmt.Println(RemoveData([]int{1, 1, 2, 2, 2, 6, 6, 8})) // [1 2 6 8]
}
