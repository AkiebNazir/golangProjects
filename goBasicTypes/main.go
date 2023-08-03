package main

import (
	"fmt"

	"github.com/goProjects/goBasicTypes/parseJson"
	"github.com/goProjects/goBasicTypes/stringGo"
	"golang.org/x/example/hello/reverse"
)

func main() {
	fmt.Println("module basic types")
	fmt.Println(reverse.String("hello"))
	stringGo.CustomerStrings(20, 400)

	// parsing json data from a file

	fmt.Printf("%+v\n", parseJson.ParseJsonData("./parseJson/data.json"))

}
