package main

import (
	"math/rand"
	"net/http"
	"text/template"
	"time"
	"fmt"
)

var (
	version   string
	timestamp string
)

func main() {
	http.HandleFunc("/", bOperator)
	fmt.Println("Starting go-bastard-operator on http://127.0.0.1:8080")
	http.ListenAndServe(":8080", nil)
}

func bOperator(w http.ResponseWriter, req *http.Request) {
	rnd := random()
	tmpl := template.Must(template.ParseFiles("bofh.tmpl"))
	tmpl.Execute(w, struct {
		Version   string
		TimeStamp string
		Excuse    string
	}{
		Version:   version,
		TimeStamp: timestamp,
		Excuse:    excuses[rnd],
	})
}

func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(len(excuses) - 1)
}
