package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	host     = "192.168.1.174"
	port     = 5432
	user     = "postgres"
	password = "9PstG1s>N(50zo3v"
	dbname   = "frag-herakles"
)

func main() {
	rq()
}

func handleAnswer(w http.ResponseWriter, r *http.Request) {
	fullPath := r.URL.Path
	remainingPath := fullPath[len("/answer/"):] // Get the portion after "/answer/"
	// Connect with Database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// Search for the word in the database
	query := "SELECT * FROM vocabulary WHERE greek = '" + remainingPath + "'"
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if !rows.Next() {
		fmt.Println("Nichts Reihe")
		return
	}

	var (
		col1 interface{}
		col2 interface{}
	)

	err = rows.Scan(&col1, &col2)
	if err != nil {
		panic(err)
	}

	fmt.Fprint(w, col1, "	", col2)
}

func rq() {
	http.HandleFunc("/answer/", handleAnswer)
	fmt.Println("Server listening on Port 8080")
	http.ListenAndServe(":8080", nil)
}
