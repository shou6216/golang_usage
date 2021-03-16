package db

import (
	// ビルド時にコンパイルの必要があるので_使ってimport
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const DbPath = "assets/usage.sql"

var DbConnection *sql.DB

type Deposit struct {
	Date  string
	Money int
}

func init() {
	log.Println("init sqlite start")

	DbConnection, _ := sql.Open("sqlite3", DbPath)
	defer DbConnection.Close()
	cmd := `CREATE TABLE IF NOT EXISTS deposit(
		date STRING PRIMARY KEY,
		money INT)`
	_, err := DbConnection.Exec(cmd)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("init sqlite end")
}

func SaveOrUpdate(date string, money int) {
	DbConnection, _ := sql.Open("sqlite3", DbPath)
	defer DbConnection.Close()
	cmd := "REPLACE INTO deposit (date, money) VALUES (?, ?)"
	_, err := DbConnection.Exec(cmd, date, money)
	if err != nil {
		log.Fatal(err)
	}
}

func FindAll() []Deposit {
	DbConnection, _ := sql.Open("sqlite3", DbPath)
	defer DbConnection.Close()
	cmd := "SELECT * FROM deposit"
	rows, _ := DbConnection.Query(cmd)
	defer rows.Close()
	var deposits []Deposit
	for rows.Next() {
		var deposit Deposit
		err := rows.Scan(&deposit.Date, &deposit.Money)
		if err != nil {
			log.Println(err)
		}
		deposits = append(deposits, deposit)
	}

	err := rows.Err()
	if err != nil {
		log.Fatalln(err)
	}

	return deposits
}
