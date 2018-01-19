package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("vim-go")
	db, err := sql.Open("postgres", "user=mike password=19850922 dbname=mike sslmode=disable")
	if err == nil {
		fmt.Println(db)
	}
	rows, _ := db.Query("SELECT * FROM mytable")
	fmt.Println(rows)
	for rows.Next() {
		var uid int
		var username string
		rows.Scan(&uid, &username)
		fmt.Println(uid)
		fmt.Println(username)
	}
}
