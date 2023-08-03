package main

import (
	"fmt"

	"github.com/goProjects/goBasicTypes/parseJson"
)

func main() {
	fmt.Println("module basic types")

	// parsing json data from a file

	fmt.Printf("%+v\n", parseJson.ParseJsonData("./parseJson/data.json"))

}
