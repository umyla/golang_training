package main

import (
	"errors"
	"fmt"
	"strconv"
)

var ErrNoRecordFound = errors.New("Record Not found")
var ErrNotMatch = errors.New("Unable to found matching record")
var user = map[int]string{
	1: "raj",
}

type ErrorFound struct {
	Func  string
	Input string
	Err   error
}

func (a *ErrorFound) Error() string {
	return "main." + a.Func + "enterd" + a.Input + ": " + a.Err.Error()
}
func searchrecord(id int) error {
	name, ok := user[id]
	if !ok {
		return &ErrorFound{
			Func:  "searchrecord",
			Input: strconv.Itoa(id),
			Err:   ErrNoRecordFound,
		}
	}

	fmt.Println(name)
	return nil
}
func otherfunc(data string) error {
	return &ErrorFound{
		Func:  "Newfunc",
		Input: data,
		Err:   ErrNotMatch,
	}
}

func main() {
	err := searchrecord(10)
	fmt.Println(err)
	err = otherfunc("hello")
	fmt.Println(err)
	_, err = strconv.Atoi("abc")
	//os.PathError{}

}
