package main

import (
	"bufio"
	"io"
	"strings"
)

func cl(r io.Reader) int {
	s := bufio.NewScanner(r)

	n := 0

	for s.Scan() {
		if len(strings.TrimSpace(s.Text())) > 0 {
			n += 1
		}
	}

	return n
}
