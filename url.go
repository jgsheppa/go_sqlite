package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/pkg/browser"
)

func openBrowser(db *sql.DB) {
	var url string
	var id int
	
	rows, err := db.Query(`SELECT * FROM urls;`)
	if err != nil {
    log.Fatal(err)
  }
  defer rows.Close()
  for rows.Next() {
		err =	rows.Scan(
				&id, &url,
			)
		if err != nil {
			log.Fatal(err)
		}
		browser.OpenURL(url)
	}
}

func insertURL(db *sql.DB, url string) {
	stmt, err := db.Prepare(`INSERT INTO urls(url) VALUES(?)`)
	if err != nil {
    log.Fatal(err)
  }
	if _, err := stmt.Exec(url); err != nil {
		log.Fatal(err)
	}
	fmt.Println("URL successfully added to database")
}