package main

func main() {
	var arr []int = []int{1, 2, 3}
	for idx, value := range arr {
		println(idx, value)
	}

	m := make(map[int]int)
	for i := 0; i < 1000000; i++ {
		m[i] = i * 2
	}

	v := m[20]
	println("completed", v)
}
