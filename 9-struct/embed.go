package main

import "fmt"

type movies struct {
	actor, actress string
}

func (m *movies) updateactor(actor string) {
	m.actor = actor
}

type horror struct {
	movies
	director string
}

func (h horror) show() {
	fmt.Println((h))
}

type comedy struct {
	movies
	budget []int
}

func (c comedy) show() {
	fmt.Println(c)
}

func main() {
	a := horror{
		movies: movies{
			actor:   "hero1",
			actress: "heroine",
		},
		director: "dop",
	}
	a.show()
	a.updateactor("star")
	a.show()
	a.updateactor("sta1r")
	a.show()
	b := comedy{
		movies: movies{
			actor:   "hero2",
			actress: "heroine2",
		},
		budget: []int{1, 2},
	}
	b.show()
	b.updateactor("superstar")
	b.show()
}
