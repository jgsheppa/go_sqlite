package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/pkg/browser"
)

func (db *DB) Start() {
	runPrograms(db.SQLite)
	// Allow time for Firefox to boot before
	// opening URLs in browser
	time.Sleep(3 * time.Second)

	var url string
	var id int

	rows, err := db.SQLite.Query(`SELECT * FROM urls;`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id, &url)
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

func (db *DB) InsertURL(url string) {
	stmt, err := db.SQLite.Prepare(`INSERT INTO urls(url) VALUES(?)`)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := stmt.Exec(url); err != nil {
		log.Fatal(err)
	}
	fmt.Println("URL successfully added to database")
}

func (db *DB) DeleteURL(id int) {
	stmt, err := db.SQLite.Prepare(`DELETE FROM urls WHERE id = ?`)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := stmt.Exec(id); err != nil {
		log.Fatal(err)
	}
	fmt.Println("URL successfully deleted from database")
}

func (db *DB) ListURLs() {
	var url string
	var id int

	rows, err := db.SQLite.Query(`SELECT * FROM urls;`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&id, &url,
		)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d: %s\n", id, url)
	}
}

func readUserURL() (string, string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter URL ID: ")
	id, err := reader.ReadString('\n')
	id = strings.Trim(id, "\n")

	if err != nil {
		return "", "", err
	}
	fmt.Print("Enter updated URL: ")
	url, err := reader.ReadString('\n')
	url = strings.Trim(url, "\n")
	if err != nil {
		return "", "", err
	}

	return id, url, nil
}

func (db *DB) UpdateURL() {
	id, url, err := readUserURL()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)
	fmt.Println(url)

	stmt, err := db.SQLite.Prepare(`UPDATE urls SET url = ? WHERE id = ?`)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := stmt.Exec(url, id); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Program successfully updated")
}
