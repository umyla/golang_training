package main

import (
	"fmt"
)

type reader interface {
	read() int
}
type writer interface {
	write(data string)
}
type readWriter interface {
	reader
	writer
}
type user struct {
	name string
}

func (u user) read() int {
	return 10
}
func (u user) write(data string) {
	fmt.Println(data)
}
func abc(rw readWriter) {
	fmt.Println(rw.read())
	rw.write("hello")

}
func main() {
	var u user
	abc(u)
	fmt.Println("readWrite implemented")
}
