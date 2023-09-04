package classes

import (
	"database/sql"
	"fmt"
)

func Add(db *sql.DB, name string, teacher string, room string, period int) error {
	if db == nil {
		panic("db is nil")
	}
	_, err := db.Exec("INSERT INTO classes (name, teacher, room, period) VALUES (?, ?, ?, ?)", name, teacher, room, period)
	if err != nil {
		return err
	}
	return nil
}

func GetAll(db *sql.DB) {
	if db == nil {
		panic("db is nil")
	}
	rows, err := db.Query("SELECT * FROM classes ORDER BY PERIOD ASC")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	fmt.Println()
	for rows.Next() {
		var id int
		var name string
		var teacher string
		var room string
		var period int
		err = rows.Scan(&id, &name, &teacher, &room, &period)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("%s, %s, ROOM %s, PERIOD %d\n", name, teacher, room, period)
	}
	fmt.Println()
	err = rows.Err()
	if err != nil {
		panic(err.Error())
	}
}
