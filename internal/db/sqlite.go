package db

import (
	// ビルド時にコンパイルの必要があるので_使ってimport
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DbConnection *sql.DB

type Person3 struct {
	Name string
	Age  int
}

func init() {
	log.Println("init sqlite start")
	log.Println("init sqlite end")
}

func Database() {
	// https://github.com/mattn/go-sqlite3
	// https://github.com/mattn/go-sqlite3#installation
	// gcc必要
	DbConnection, _ := sql.Open("sqlite3", "./sql/example.sql")
	defer DbConnection.Close()
	cmd := `CREATE TABLE IF NOT EXISTS person(
		name STRING,
		age INT)`
	_, err := DbConnection.Exec(cmd)
	if err != nil {
		log.Fatal(err)
	}

	cmd = "INSERT INTO person (name, age) VALUES (?, ?)"
	_, err = DbConnection.Exec(cmd, "Mike", 25)
	if err != nil {
		log.Fatal(err)
	}

	// multiple select
	cmd = "SELECT * FROM person"
	rows, _ := DbConnection.Query(cmd)
	defer rows.Close()
	var pp []Person3
	for rows.Next() {
		var p Person3
		err := rows.Scan(&p.Name, &p.Age)
		if err != nil {
			log.Println(err)
		}
		pp = append(pp, p)
	}
	err = rows.Err()
	if err != nil {
		log.Fatalln(err)
	}
	for _, p := range pp {
		fmt.Println(p.Name, p.Age)
	}

	// single select
	cmd = "SELECT * FROM person where age = ?"
	row := DbConnection.QueryRow(cmd, 20)
	var p Person3
	err = row.Scan(&p.Name, &p.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No row")
		} else {
			log.Println(err)
		}
	}
	fmt.Println(p.Name, p.Age)
}
