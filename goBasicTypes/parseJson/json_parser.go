package parseJson

import (
	"encoding/json"
	"fmt"
	"os"
)

type JsonData struct {
	Person []Person `json:"person"`
}

type Person struct {
	Name  string `json:"name" validate:"required"`
	Age   int    `json:"age" validate:"gte=18,lte=30"`
	Email string `json:"email"`
}

func ParseJsonData(file_name string) JsonData {
	data, err := os.ReadFile(file_name)

	if err != nil {
		fmt.Println("error while reading the file ", err)
		os.Exit(1)
	}

	var jsondata JsonData
	err = json.Unmarshal(data, &jsondata)
	if err != nil {
		fmt.Println("Error while parsing into struct ", err)
		os.Exit(1)
	}
	return jsondata

}
