package main

import (
	"fmt"
	"net/http"
)

func handleAnswer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Answer")
}

func main() {
	http.HandleFunc("/answer", handleAnswer)
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
