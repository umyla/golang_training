package main

import (
	"fmt"
)

type transport interface {
	capacity() int
}

type bus struct {
	SeatsInRow int
	rows       int
}

func (b bus) capacity() int {
	s := b.rows * b.SeatsInRow
	fmt.Println("total capacity of bus", s)
	return s
}

type car struct {
	seats int
	rows  int
}

func (c car) capacity() int {
	a := c.rows * c.seats
	fmt.Println("total capacity of car", a)
	return a
}
func Fetch(t transport) int {
	return t.capacity()
}
func main() {
	s1 := bus{3, 5}
	a1 := car{2, 3}

	Fetch(s1)
	Fetch(a1)

}
