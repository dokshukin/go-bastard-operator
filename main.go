package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var (
	Version   string
	Timestamp string
)

func main() {

	fmt.Println("Version:\t", Version)
	fmt.Println("build.Time:\t", Timestamp)

	http.HandleFunc("/", bOperator)
	http.ListenAndServe(":8080", nil)
}

func bOperator(w http.ResponseWriter, req *http.Request) {
	rnd := random()
	reply := getReply(rnd)
	fmt.Fprintf(w, reply)
}

func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(len(excuses) - 1)
}

func getReply(i int) (reply string) {
	reply = `<!DOCTYPE html>
		<html>
		<head>
		<title>Lame excuses from bastard operator</title>
		<meta charset=\"utf-8\">
		</head>
		<body>
    <center>
    <h6>bastard operator version ` + Version +
		" (build: " + Timestamp + `)</h6>
		<h2>The core issue is:</2>
		<h1>` + excuses[i] + `</1>
		</center>
		</body>
    </html>`

	return reply
}
