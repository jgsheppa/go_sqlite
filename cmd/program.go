package cmd

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
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
		err =	rows.Scan(&id, &program,)
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

func programByID(db *sql.DB, program string) (int, error) {	
	var id int
	
	rows, err := db.Query(`SELECT FROM programs WHERE url = ?;`, program)
	if err != nil {
    return 0, err
  }
  defer rows.Close()

  for rows.Next() {
		err =	rows.Scan(
				&id, &program,
			)
		if err != nil {
			return 0, err
		}
		fmt.Println(id)
		return id, nil
	}
	return 0, nil
}

func InsertProgram(db *sql.DB, program string) {
	stmt, err := db.Prepare(`INSERT INTO programs(program) VALUES(?)`)
	if err != nil {
    log.Fatal(err)
  }
	if _, err := stmt.Exec(program); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Program successfully added to database")
}

func DeleteProgram(db *sql.DB, id int) {
	stmt, err := db.Prepare(`DELETE FROM programs WHERE id = ?`)
	if err != nil {
    log.Fatal(err)
  }
	if _, err := stmt.Exec(id); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Program successfully deleted from database")
}

func ListPrograms(db *sql.DB) {
	var program string
	var id int
	
	rows, err := db.Query(`SELECT * FROM programs;`)
	if err != nil {
    log.Fatal(err)
  }
  defer rows.Close()

  for rows.Next() {
		err =	rows.Scan(
				&id, &program,
			)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d: %s\n",id, program)
	}
}