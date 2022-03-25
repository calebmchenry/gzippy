package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/NYTimes/gziphandler"
)

type db struct {
	Ten    []byte
	Twenty []byte
}

func (db *db) load() {
	data, err := os.ReadFile("static/20MB.json")
	if err != nil {
		panic("oh no!")
	}
	db.Twenty = data

	data, err = os.ReadFile("static/10MB.json")
	if err != nil {
		panic("oh no!")
	}
	db.Ten = data
}

func main() {
	db := db{}
	db.load()

	tenMB := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(db.Ten)
	})
	twentyMB := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(db.Twenty)
	})

	gzipTenMB := gziphandler.GzipHandler(tenMB)
	gzipTwentyMB := gziphandler.GzipHandler(twentyMB)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.Handle("/gzip/10MB", gzipTenMB)
	http.Handle("/gzip/20MB", gzipTwentyMB)
	http.Handle("/raw/10MB", tenMB)
	http.Handle("/raw/20MB", twentyMB)
	fmt.Println("Running on localhost:8000")
	fmt.Println(http.ListenAndServe("localhost:8000", nil))
}
