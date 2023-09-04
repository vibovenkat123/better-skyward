package main

import (
	database "better-skyward/pkg/db"
	"fmt"
	"log"
)

func main() {
	db, err := database.Open()
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer db.Close()
	fmt.Print(`
Better Skyward
--------------
Choose an option:
1) Add a class
2) List schedule
> `)
	var choice int
	_, err = fmt.Scanln(&choice)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
