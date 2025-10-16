package main

import "fmt"

func main() {
	var arr [3]int = [3]int{1, 2, 3}

	slice := []int{1, 2, 3}
	slice = append(slice, 4)

	fmt.Println("Array", arr)
	fmt.Println("slice: ", slice)

}
