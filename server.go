package main

import (
	"fmt"

	"example.com/user/hello/db"
	"example.com/user/hello/route"
)

func main() {
	db.Init()

	e := route.Init()
	fmt.Printf("000111")
	e.Logger.Fatal(e.Start(":1323"))

}
