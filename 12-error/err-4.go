package main

import (
	"fmt"
	"strconv"
)

func main() {
	_, err := strconv.Atoi("abc")
	fmt.Println(err)
	_, err = strconv.Atoi("xyz")
	fmt.Println(err)
	_, err = strconv.ParseInt("abc", 10, 64)
	fmt.Println(err)
}
