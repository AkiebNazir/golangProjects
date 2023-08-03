package main

import (
	"fmt"
	"ormgorm/handler_func"
)

func main() {
	fmt.Println("Connecting to a database ..... !")
	db, err := handler_func.ConnectoToDB()
	if err != nil {
		fmt.Println("database connection error: ", err)
	}
	app := handler_func.App{
		Db: db,
	}
	data, err := app.CreateTable()
	if err != nil {
		fmt.Println("Some Kind of Error: ", err)
		return
	}

	fmt.Println(data)

}
