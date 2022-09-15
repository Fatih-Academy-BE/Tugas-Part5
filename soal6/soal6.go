package main

import "fmt"

func MeanMedian(arrayInput []float64) (float64, float64) {
	// your code here

}

func main() {
	fmt.Println(MeanMedian([]float64{1, 2, 3, 4}))                          // 2.5 2.5
	fmt.Println(MeanMedian([]float64{1, 3, 5, 7, 11}))                      // 5.4 5
	fmt.Println(MeanMedian([]float64{2, 4, 6, 8, 10}))                      // 6 6
	fmt.Println(MeanMedian([]float64{10, 20, 30, 40, 50}))                  // 30 30
	fmt.Println(MeanMedian([]float64{25, 50, 75, 100, 125, 150, 175, 200})) // 112.5 112.5
}
