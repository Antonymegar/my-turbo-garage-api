package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"

	_ "github.com/lib/pq"
)

func main() {

	conn, _ := url.Parse()
	conn.RawQuery = "sslmode=verify-ca;sslrootcert=ca.pem"

	db, err := sql.Open("postgres", conn.String())

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT version()")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var result string
		err = rows.Scan(&result)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Version: %s\n", result)
	}
}