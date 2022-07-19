package main

type movies struct {
	actor, actress string
}
type hit struct {
	m      movies
	rating []int
}

func main() {
	a := hit{
		movies: movies{
			actor:   "jay",
			actress: "anu",
		},
		rating: []int{10},
	}
}
