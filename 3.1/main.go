package main

import "strings"

func main() {
	s := strings.ToUpper("dominic")
	s = strings.Replace(s, "O", "o", -1)
	println(s)
}
