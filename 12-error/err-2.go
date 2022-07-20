package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var ErrNoRecord = errors.New("There is no record")

func main() {
	f, err := OpenFile("text.json")
	if err != nil {
		log.Fatalln(err)
	}
	f.Close()
}
func OpenFile(filename string) (*os.File, error) {
	f, err := os.OpenFile(filename, os.O_WRONLY, 0666)
	if err != nil {
		return nil, fmt.Errorf("custom msg %v %v", err, ErrNoRecord)
	}
	return f, nil
}
