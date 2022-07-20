package main

import (
	"errors"
	"fmt"
	"log"
)

var student = make(map[int]string, 1000)
var ErrRecord = errors.New("record not found")

func fetchrecord(id int) (string, error) {
	name, ok := student[id]
	if !ok {
		return "", ErrRecord
	}
	return name, nil
}
func main() {
	data, err := fetchrecord(100)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(data)
}
