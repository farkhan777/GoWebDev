package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
)

type Amigos struct {
	aID   int    `json:"aID"`
	aName string `json:"aName"`
}

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "root:root@/test02")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":7777", nil)
	checkError(err)
}

func index(w http.ResponseWriter, r *http.Request) {
	rows, _ := db.Query("SELECT * FROM Amigos WHERE aID = 1")
	defer rows.Close()
	for rows.Next() {
		var a Amigos
		err = rows.Scan(&a.aID, &a.aName)
		checkError(err)
		myName := a.aName
		io.WriteString(w, myName+"\n")
	}
	_, err := io.WriteString(w, "\nSuccessfully completed.")
	if err != nil {
		panic(err)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
