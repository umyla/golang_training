package main

import "fmt"

type class struct {
	name, subject string
}

var studentId = map[int]class{
	1: {
		name:    "abhi",
		subject: "english",
	},
	2: {
		name:    "ben",
		subject: "maths",
	},
}

func main() {
	key := 2
	fmt.Println(studentId[key])
	searchRecord(4)
}
func searchRecord(id int) {
	u, ok := studentId[id]
	if !ok {
		fmt.Println("studentId not found", ok)
	}
	return

	fmt.Println(u)
}
