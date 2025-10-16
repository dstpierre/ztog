package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		total := calcAir(i)
		if total%2 == 0 {
			fmt.Println(i, " is even")
		} else {
			fmt.Println(i, " is odd")
		}
	}

}

func calcAir(i int) int {
	return i * i
}
