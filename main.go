package main

import (
	"database/sql"
	"flag"
	"log"
	"strings"

	"github.com/jgsheppa/go_sqlite/cmd"
	_ "github.com/mattn/go-sqlite3"
)



func main() {

	db, err := sql.Open("sqlite3", "shakespeare.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var line, add string
	var id int
	var isStarted, isUrl, isProgram, list bool

	flag.StringVar(&line, "line", "", "Player line")
	flag.BoolVar(&isStarted, "start", false, "Starts programs")
	flag.BoolVar(&isUrl, "url", false, "Flag for url actions")
	flag.BoolVar(&isProgram, "program", false, "Flag for program actions")
	flag.BoolVar(&list, "list", false, "List all urls")
	flag.StringVar(&add, "add", "", "Adds new url or program")
	flag.IntVar(&id, "delete", 0, "Deletes url or program")
	flag.Parse()
	
	switch true {
	case isStarted:
		cmd.Start(db)
	case isUrl && len(strings.TrimSpace(line)) > 0:
		search(db, line)
	case isUrl && list:
		cmd.ListURLs(db)
	case isUrl && len(strings.TrimSpace(add)) > 0:
		cmd.InsertURL(db, add)
	case isUrl && id != 0:
		cmd.DeleteURL(db, id)
	case isProgram && list:
		cmd.ListPrograms(db)
	case isProgram && len(strings.TrimSpace(add)) > 0:
		cmd.InsertProgram(db, add)
	case isProgram && id != 0:
		cmd.DeleteProgram(db, id)
	}
}

