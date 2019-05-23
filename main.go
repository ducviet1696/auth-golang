package main

import (
	"authGoLang/router"
	"fmt"
)

func main() {
	fmt.Println("Welcome to the webserver")
	e := router.New()

	_ = e.Start(":8000")
}
