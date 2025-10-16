package main

import (
	"errors"
	"fmt"
	"os"
)

// fileExists checks if a file exists and is not a directory.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if err == nil {
		return !info.IsDir()
	}

	if errors.Is(err, os.ErrNotExist) {
		return false
	}

	fmt.Printf("Error accessing file %s: %v\n", filename, err)
	return false
}
