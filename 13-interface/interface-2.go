package main

import (
	"fmt"
)

type homes interface {
	property() int
}

type SingleHome struct {
	rooms   int
	cost    int
	intrest float64
}

func (s SingleHome) property() int {
	h1 := s
	//fmt.Println(h1)
	return h1.cost
}

type TownHome struct {
	SingleHome
	location string
}

func (t TownHome) property() int {
	var h2 = t
	fmt.Println(h2)
	return t.rooms
}
func realestate(h homes) int {
	fmt.Println("test")
	return h.property()

}

func main() {
	s1 := SingleHome{
		rooms:   4,
		cost:    100000,
		intrest: 5.2,
	}
	t1 := TownHome{
		SingleHome: SingleHome{
			rooms:   3,
			cost:    50000,
			intrest: 5.6,
		},
		location: "Ind",
	}
	//fmt.Println(s1)
	//fmt.Println(t1)
	cost := realestate(s1)
	fmt.Println(cost)
	rooms := realestate(t1)
	fmt.Println(rooms)
}
