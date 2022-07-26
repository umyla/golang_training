package main

import "fmt"

type Walker interface {
	walk()
}
type Runner interface {
	run()
}
type WalkRunner interface {
	Walker
	Runner
}
type Human struct {
	name string
}

func (h Human) walk() {

}
func (h Human) run() {

}

type Robo struct {
	name string
}

func (r Robo) walk() {

}

func main() {
	h := Human{"Sam"}
	r := Robo{"xyz"}
	_, _ = h, r
	var a Walker

	x, ok := a.(Human)
	if !ok {
		fmt.Println("r is not storing the instance of the human struct")
		fmt.Println("you are not robo I can't hire you")
		return
	}
	fmt.Println("yes you are a human", x.name)
}
