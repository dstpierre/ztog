package main

import "strconv"

func main() {
	i, err := strconv.ParseInt("123", 10, 64)
	if err != nil {
		println(err)
		return
	}
	println(i)
}
