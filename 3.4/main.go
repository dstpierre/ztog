package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func main() {
	s := time.Now().Format("Monday 05")

	values := strings.Split(s, " ")

	second, err := strconv.Atoi(values[1])
	if err != nil {
		log.Fatal(err, s)
	}

	var i int
	for i = 0; i < 59; i++ {
		if i == second {
			break
		}
	}

	fmt.Printf("It took %d number of loop to reach %d\n", i, second)
	println("Today is: ", values[0])
}
