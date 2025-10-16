package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("trx.dat")
	if err != nil {
		fmt.Println("cannot create transaction file:", err)
		os.Exit(1)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("unable to read from file:", err)
		os.Exit(1)
	}

	if err = os.WriteFile("trx2.dat", content, 0644); err != nil {
		fmt.Println("unable to write data to file:", err)
		os.Exit(1)
	}
}
