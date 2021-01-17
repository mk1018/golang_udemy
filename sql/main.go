package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {

	insert()
	// select1()
	select2()

}

func create() {
	Db, _ := sql.Open("sqlite3", "./example.sql")

	defer Db.Close()

	cmd := `CREATE TABLE IF NOT EXISTS persons(
				name STRING,
				age INT)`

	_, err := Db.Exec(cmd)

	if err != nil {
		log.Fatalln(err)
	}
}

func insert() {
	Db, _ := sql.Open("sqlite3", "./example.sql")

	defer Db.Close()

	cmd := `INSERT INTO persons(name, age) VALUES(?, ?)`

	_, err := Db.Exec(cmd, "tarou", 20)

	if err != nil {
		log.Fatalln(err)
	}
}

func select1() {
	Db, _ := sql.Open("sqlite3", "./example.sql")
	defer Db.Close()

	cmd := "SELECT * FROM persons where age = ?"

	row := Db.QueryRow(cmd, 20)
	var p Person
	err := row.Scan(&p.Name, &p.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No row")
		} else {
			log.Println(err)
		}
	}

	fmt.Println(p.Name, p.Age)
}

func select2() {
	Db, _ := sql.Open("sqlite3", "./example.sql")
	defer Db.Close()

	cmd := "SELECT * FROM persons"

	rows, _ := Db.Query(cmd)
	defer rows.Close()
	var pp []Person

	for rows.Next() {
		var p Person
		err := rows.Scan(&p.Name, &p.Age)
		if err != nil {
			log.Println(err)
		}
		pp = append(pp, p)
	}

	for _, p := range pp {
		fmt.Println(p)
	}

}
