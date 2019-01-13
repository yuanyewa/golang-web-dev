package main

import (
	"fmt"
	"net/http"
)

type mytype string

func (m mytype) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is yuanye's type")
}

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Any code you want in this func")
}

func main() {
	var d hotdog
	go http.ListenAndServe(":8080", d)
	var m mytype
	go http.ListenAndServe(":1234", m)
	for {

	}
}
