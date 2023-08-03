package main

import (
	"fmt"

	"github.com/goProjects/goBasicTypes/arrayslices"
	"github.com/goProjects/goBasicTypes/parseJson"
)

func main() {
	fmt.Println("module basic types")

	// parsing json data from a file

	fmt.Printf("%+v\n", parseJson.ParseJsonData("./parseJson/data.json"))

	// array and slices in god

	var arr arrayslices.ArrayNums
	var sl arrayslices.SliceNums
	for i := 1; i < 21; i++ {
		arr.Arr[i-1] = i
		sl.Arr = append(sl.Arr, i)
	}
	arr.ArrayMonupulation()
	sl.SliceEdit()

	fmt.Printf("the array elements are :%v\n", arr)
	fmt.Printf("the slice elements are :%v\n", sl)

}
