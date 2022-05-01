package main

import (
	"database/sql"
	"flag"
	"log"
	"strings"

	"github.com/jgsheppa/go_sqlite/cmd"
	_ "github.com/mattn/go-sqlite3"
)

func StartDB() (cmd.DB, error) {
	db, err := sql.Open("sqlite3", "shakespeare.db")
	if err != nil {
		return cmd.DB{}, err
	}

	return cmd.DB{
		SQLite: db,
	}, nil
}

func main() {

	db, err := StartDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQLite.Close()

	var line, add string
	var id int
	var isStarted, isUrl, isProgram, list, update bool

	flag.StringVar(&line, "line", "", "Player line")
	flag.BoolVar(&isStarted, "start", false, "Starts programs")
	flag.BoolVar(&isUrl, "url", false, "Flag for url actions")
	flag.BoolVar(&isProgram, "program", false, "Flag for program actions")
	flag.BoolVar(&list, "list", false, "List all urls")
	flag.BoolVar(&update, "update", false, "Update program")
	flag.StringVar(&add, "add", "", "Adds new url or program")
	flag.IntVar(&id, "delete", 0, "Deletes url or program")
	flag.Parse()

	switch true {
	case isStarted:
		db.Start()
	case isUrl && list:
		db.ListURLs()
	case isUrl && len(strings.TrimSpace(add)) > 0:
		db.InsertURL(add)
	case isUrl && id != 0:
		db.DeleteURL(id)
	case isUrl && update:
		db.UpdateURL()
	case isProgram && list:
		db.ListPrograms()
	case isProgram && len(strings.TrimSpace(add)) > 0:
		db.InsertProgram(add)
	case isProgram && id != 0:
		db.DeleteProgram(id)
	case isProgram && update:
		db.UpdateProgram()
	}
}
