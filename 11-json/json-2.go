package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type movies struct {
	Actor      string `json."actor"`
	Actress    string `json."actress"`
	TotalFilms int    `json."total_films`
}

// type horror struct {
// 	movies
// 	director string
// }

// func (h horror) show() {
// 	fmt.Println((h))
// }

// type comedy struct {
// 	movies
// 	budget []int
// }

// func (c comedy) show() {
// 	fmt.Println(c)
// }

func main() {
	cinema := []movies{
		{
			Actor:      "hero1",
			Actress:    "heroine",
			TotalFilms: 3,
		},
		{
			Actor:      "hero2",
			Actress:    "heroine2",
			TotalFilms: 5,
		},
	}

	jsonData, err := json.MarshalIndent(cinema, "", "\t")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(jsonData))
	f, err := os.OpenFile("movies.json", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	_, err = f.Write(jsonData)
	if err != nil {
		log.Fatalln(err)
	}
}
