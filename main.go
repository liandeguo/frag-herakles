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

func handleAnswer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Answer")
}

func main() {

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

	fmt.Println("Successfully connected!")
	http.HandleFunc("/answer", handleAnswer)
	http.ListenAndServe(":8080", nil)
}
