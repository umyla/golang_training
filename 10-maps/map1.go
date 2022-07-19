package main

import "fmt"

func main() {
	class := map[string]string{
		"name":      "student1",
		"subject":   "maths",
		"studentId": "123",
	}
	fmt.Println(class["name"])
	class["subject"] = "science"
	class["score"] = "500"
	fmt.Println(class)
}
