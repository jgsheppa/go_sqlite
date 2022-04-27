package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)



func main() {

	db, err := sql.Open("sqlite3", "shakespeare.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var line, newURL string
	var urls bool

	flag.StringVar(&line, "line", "", "Player line")
	flag.BoolVar(&urls, "url", false, "Open urls in browser")
	flag.StringVar(&newURL, "add", "", "Adds new url to database")
	flag.Parse()
	
	switch true {
	case len(strings.TrimSpace(line)) > 0:
		search(db, line)
		fmt.Print("1")
	case urls:
		openBrowser(db)
	case len(strings.TrimSpace(newURL)) > 0:
		insertURL(db, newURL)
	}

	
	if err != nil {
		log.Fatal(err)
	}
	
}