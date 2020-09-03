package main

import (
	"math/rand"
	"net/http"
	"text/template"
	"time"
)

var (
	version   string
	timestamp string
)

func main() {
	http.HandleFunc("/", bOperator)
	http.ListenAndServe(":8080", nil)
}

func bOperator(w http.ResponseWriter, req *http.Request) {
	rnd := random()
	tmpl := template.Must(template.ParseFiles("./templates/bofh.html"))
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
