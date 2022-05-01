package cmd

import (
	"bufio"
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

// Stdout is the io.Writer to which executed commands write standard output.
var Stdout io.Writer = os.Stdout

// Stderr is the io.Writer to which executed commands write standard error.
var Stderr io.Writer = os.Stderr

func runProgram(program string) error {
	cmd := exec.Command("open", "-a", program)
	cmd.Stdout = Stdout
	cmd.Stderr = Stderr
	return cmd.Run()
}

func runPrograms(db *sql.DB) {
	var program string
	var id int

	rows, err := db.Query(`SELECT * FROM programs;`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id, &program)
		if err != nil {
			log.Fatal(err)
		}
		err := runProgram(program)
		if err != nil {
			log.Fatal(err)
		}

	}

	err = runProgram("Visual Studio Code")
	if err != nil {
		log.Fatal(err)
	}
	err = runProgram("Firefox")
	if err != nil {
		log.Fatal(err)
	}
}

func (db *DB) InsertProgram(program string) {
	stmt, err := db.SQLite.Prepare(`INSERT INTO programs(program) VALUES(?)`)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := stmt.Exec(program); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Program successfully added to database")
}

func readUserInput() (string, string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter program ID: ")
	id, err := reader.ReadString('\n')
	id = strings.Trim(id, "\n")
	if err != nil {
		return "", "", err
	}
	fmt.Print("Enter updated program name: ")
	program, err := reader.ReadString('\n')
	program = strings.Trim(program, "\n")
	if err != nil {
		return "", "", err
	}

	return id, program, nil
}

func (db *DB) UpdateProgram() {
	id, program, err := readUserInput()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.SQLite.Prepare(`UPDATE programs SET program = ? WHERE id = ?`)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := stmt.Exec(program, id); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Program successfully updated")
}

func (db *DB) DeleteProgram(id int) {
	stmt, err := db.SQLite.Prepare(`DELETE FROM programs WHERE id = ?`)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := stmt.Exec(id); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Program successfully deleted from database")
}

func (db *DB) ListPrograms() {
	var program string
	var id int

	rows, err := db.SQLite.Query(`SELECT * FROM programs;`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&id, &program,
		)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d: %s\n", id, program)
	}
}
