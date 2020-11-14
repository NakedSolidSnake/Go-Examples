package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Aluno struct {
	ID        int
	Nome      string
	Matricula int64
	Idade     int64
	Curso     string
}

func main() {
	// Open up our database connection.
	db, err := sql.Open("mysql", "root:$0l1d$2r2@tcp(127.0.0.1:3306)/aluno")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT * FROM alunos")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var aluno Aluno
		// for each row, scan the result into our tag composite object
		err = results.Scan(&aluno.ID, &aluno.Nome, &aluno.Matricula, &aluno.Idade, &aluno.Curso)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		fmt.Println(aluno)
	}

}
