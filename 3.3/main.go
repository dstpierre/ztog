package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	yesterday, err := time.Parse("2006-01-02 15:04", "2024-03-12 14:25")
	if err != nil {
		println(err)
		return
	}

	fmt.Println(now)
	fmt.Println(yesterday)

	yesterday = yesterday.Add(24 * time.Hour)
	fmt.Println(yesterday.Format("02 Jan 2006, 15:04"))

}
