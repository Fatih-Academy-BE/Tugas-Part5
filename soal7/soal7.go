package main

import "fmt"

func JoinArray(array1 []string, array2 []string) []string {
	// your code here

}

func main() {
	fmt.Println(JoinArray([]string{"senin", "selasa", "rabu"}, []string{"senin", "rabu", "jumat", "minggu"})) // [senin selasa rabu jumat minggu]
	fmt.Println(JoinArray([]string{"apel", "semangka", "anggur"}, []string{"anggur", "semangka", "nanas"}))   // [apel semangka anggur nanas]
}
