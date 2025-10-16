package main

import "math"

type Shape interface {
	Area() int
}

type Rectangle struct {
	Width  int
	Height int
}

func (r *Rectangle) Area() int {
	if r == nil {
		return 0
	}
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() int {
	return int(math.Pi * math.Pow(c.Radius, 2))
}
