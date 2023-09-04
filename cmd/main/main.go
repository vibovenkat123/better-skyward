package main

import (
	"better-skyward/pkg/classes"
	database "better-skyward/pkg/db"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	db, err := database.Open()
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer db.Close()
	rdr := bufio.NewReader(os.Stdin)
	for {
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
		switch choice {
		case 1:
			fmt.Print("Name: ")
			name, err := rdr.ReadString('\n')
			name = name[:len(name)-1]
			if err != nil {
				log.Fatalln(err.Error())
			}
			fmt.Print("Teacher: ")
			teacher, err := rdr.ReadString('\n')
			teacher = teacher[:len(teacher)-1]
			if err != nil {
				log.Fatalln(err.Error())
			}
			fmt.Print("Room: ")
			room, err := rdr.ReadString('\n')
			room = room[:len(room)-1]
			if err != nil {
				log.Fatalln(err.Error())
			}
			fmt.Print("Period: ")
			periodraw, err := rdr.ReadString('\n')
			if err != nil {
				log.Fatalln(err.Error())
			}
			period, err := strconv.Atoi(periodraw[:len(periodraw)-1])
			if err != nil {
				log.Fatalln(err.Error())
			}
			err = classes.Add(db, name, teacher, room, period)
			if err != nil {
				log.Fatalln(err.Error())
			}
		case 2:
			classes.GetAll(db)
		}
	}
}
