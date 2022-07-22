package main

import (
	"errors"
	"fmt"
)

var ErrNotDev = errors.New("Not found at level 1")
var ErrNotProd = errors.New("Not found at level 2")
var ErrNotAdmin = errors.New("not found at level 3")

type ErrorFound struct {
	Func  string
	Input string
	Err   error
}

func (a *ErrorFound) Error() string {
	return "main." + a.Func + "enterd" + a.Input + ": " + a.Err.Error()

}
func (e *ErrorFound) Unwrap() error { return e.Err }

func dev() error {
	err := prod()
	if err != nil {

		return &ErrorFound{
			Func:  "searchrecord",
			Input: "abc",
			Err:   ErrNotDev,
		}
		fmt.Println(err)
		fmt.Errorf("%w %v", err, ErrNotDev)
	}
	return nil
}

func prod() error {
	err := admin()
	if err != nil {
		return fmt.Errorf("%w %v", err, ErrNotProd)
	}
	return nil
}
func admin() error {
	return fmt.Errorf("%w", ErrNotAdmin)
}
func main() {
	err := prod()
	fmt.Println(err)
	err = errors.Unwrap(err)
	fmt.Println(err)
}
