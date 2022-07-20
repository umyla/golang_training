package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type user struct {
	Fullname string          `json:"full_name"`
	Perms    map[string]bool `json:"perms"`
}

func main() {
	jsonData, err := os.ReadFile("test.json")
	if err != nil {
		log.Fatalln(err)
	}
	var u []user
	err = json.Unmarshal(jsonData, &u)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(u)
	for i, v := range u {
		fmt.Println(i, v)
	}
}
