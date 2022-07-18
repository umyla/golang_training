package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("test.txt", os.O_RDONLY, 0666)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()
}
