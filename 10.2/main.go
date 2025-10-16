package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Create("file")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Fprintln(file, "Hello file")
	fmt.Fprintln(os.Stdout, "Hello terminal")
}
