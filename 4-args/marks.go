package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	marks()
	fmt.Println("main func ending")
}
func marks() {
	details := os.Args[1:]
	fmt.Println(details, "len = ", len(details))
	if len(details) < 4 {
		log.Println("please provide your name, age, total marks and grade")
		//return
	}
	name := details[0]
	ageString := details[1]
	gradeString := details[2]
	marksString := details[3]

	age, err := strconv.Atoi(ageString)
	if err != nil {
		log.Println("please provide a valid age", err)
		return
	}
	grade, err := strconv.Atoi(gradeString)
	if err != nil {
		fmt.Println("please provide your grade", err)
		return
	}
	marks, err := strconv.Atoi(marksString)
	if err != nil {
		log.Println("please provide your marks", err)
		return
	}
	fmt.Println("Congratulations!")

	fmt.Println("Name=", name)
	fmt.Println("Age=", age)
	fmt.Println("marks=", marks)
	fmt.Println("Grade=", grade)
}
