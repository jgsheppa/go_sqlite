package main

import (
	"database/sql"
	"fmt"
	"log"
)

func search(db *sql.DB, line string) {
	var play, text string
	var act, scene interface{}
	
	rows, err := db.Query(`
	SELECT play, act, scene, plays.text
	FROM playsearch
	INNER JOIN plays ON playsearch.playsrowid = plays.rowid
	WHERE playsearch.text MATCH ?;
	`, line)
	if err != nil {
    // handle this error better than this
    panic(err)
  }
  defer rows.Close()
  for rows.Next() {
		err =	rows.Scan(
				&play, &act, &scene, &text,
			)
		if err != nil {
			// handle this error
			log.Fatal(err)
		}
	
		var actAndScene string
		if len(act.(string)) > 0 && len(scene.(string)) > 0 {
			actAndScene = act.(string) +"."+scene.(string)+":"
			fmt.Printf("%s %s %q\n", play, actAndScene, text)
		}

		fmt.Printf("%s %q\n", play, text)
	}
}