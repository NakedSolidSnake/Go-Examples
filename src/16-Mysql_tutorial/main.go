package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Aluno struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	UniversityNUmber int64  `json:"universitynumber"`
	Age              int8   `json:"age"`
	Description      string `json:"description"`
}

func main() {
	fmt.Println("Go MySQL Tutorial")

	db, err := sql.Open("mysql", "root:$0l1d$2r2@tcp(127.0.0.1:3306)/aluno")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert, err := db.Query("INSERT INTO alunos (nome, matricula, idade, curso) values('Teste', 1111111, 34, 'Whatever')")
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	fmt.Println("Insert Succefully")

	results, err := db.Query("Select * from alunos")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var student Aluno
		err = results.Scan(&student.Id, &student.Name, &student.UniversityNUmber, &student.Age, &student.Description)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(student)
	}

}
