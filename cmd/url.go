package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/pkg/browser"
)

func Start(db *sql.DB) {
	runPrograms(db)
	// Allow time for Firefox to boot before
	// opening URLs in browser
	time.Sleep(3 * time.Second)

	var url string
	var id int

	rows, err := db.Query(`SELECT * FROM urls;`)
	if err != nil {
    log.Fatal(err)
  }
  defer rows.Close()

  for rows.Next() {
		err =	rows.Scan(&id, &url,)
		if err != nil {
			log.Fatal(err)
		}

		err = browser.OpenURL(url)
		if err != nil {
			log.Fatal(err)
			fmt.Println("Check this URL to make sure the full URL is present in the database")
		}
	}
}

func urlByID(db *sql.DB, url string) (int, error) {	
	var id int
	
	rows, err := db.Query(`SELECT FROM urls WHERE url = ?;`, url)
	if err != nil {
    return 0, err
  }
  defer rows.Close()

  for rows.Next() {
		err =	rows.Scan(
				&id, &url,
			)
		if err != nil {
			return 0, err
		}
		fmt.Println(id)
		return id, nil
	}
	return 0, nil
}

func InsertURL(db *sql.DB, url string) {
	stmt, err := db.Prepare(`INSERT INTO urls(url) VALUES(?)`)
	if err != nil {
    log.Fatal(err)
  }
	if _, err := stmt.Exec(url); err != nil {
		log.Fatal(err)
	}
	fmt.Println("URL successfully added to database")
}

func DeleteURL(db *sql.DB, id int) {
	stmt, err := db.Prepare(`DELETE FROM urls WHERE id = ?`)
	if err != nil {
    log.Fatal(err)
  }
	if _, err := stmt.Exec(id); err != nil {
		log.Fatal(err)
	}
	fmt.Println("URL successfully deleted from database")
}

func ListURLs(db *sql.DB) {
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
		fmt.Printf("%d: %s\n",id, url)
	}
}