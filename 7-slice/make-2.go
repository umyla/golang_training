package main

import "fmt"

func main() {
	var data = make([]int, 5, 100000) // make(type,len,cap)
	var count int
	var lastCap = cap(data)
	for i := 0; i < 100000; i++ {
		data = append(data, i)

		if lastCap != cap(data) {
			count++

			capCh := float64(cap(data)-lastCap) / float64(lastCap) * 100

			lastCap = cap(data)
			fmt.Printf("Add %p Cap %d CapCh %v\n", data, cap(data), capCh)
		}
	}
	fmt.Println(count)
}
