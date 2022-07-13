package data

import (
	"fmt"
)

var DB = "postgres" // if something can break your app if any changes happens to the global var then don't use it

const Json = "json data"

func GetData() {
	//Json = "any"
	//http.StatusOK
	fmt.Println("my prod data is in postgres so use that")
	fmt.Println("getting data from ", DB, Json)
}
