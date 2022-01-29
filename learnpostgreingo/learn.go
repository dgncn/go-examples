package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type animal struct {
	id   int
	name string
	age  int
}

func main() {
	db, err := sql.Open("postgres", "user=postgres password=12345 dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal("db connection error: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("select * from dino.animal where age > $1", 5)
	handleRows(rows, err)

	// rowWith1Id := db.QueryRow("select * from animal where id == $1", 1)

	// animal := animal{}
	// err = rowWith1Id.Scan(&animal.name)
	// if err != nil {
	// 	log.Fatal(err)
	// }

}

func handleRows(rows *sql.Rows, err error) {
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	animals := []animal{}

	for rows.Next() {
		animal := animal{}
		err := rows.Scan(&animal.id, &animal.name, &animal.age)
		if err != nil {
			log.Println(err)
			continue
		}
		animals = append(animals, animal)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(animals)
}
