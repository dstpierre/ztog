package main

import "fmt"

func calculateArea(s Shape) int {
	return s.Area()
}

func main() {
	var rec *Rectangle
	c := &Circle{Radius: 1.2}

	fmt.Printf("rec: %T, %v\n", rec, rec)
	fmt.Printf("c: %T, %v\n", c, c)

	fmt.Println("rec area", calculateArea(rec))
	//fmt.Println("c area", calculateArea(c))
}
