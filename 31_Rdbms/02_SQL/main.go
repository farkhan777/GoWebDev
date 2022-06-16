package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "root:root@/test02")
	checkErr(err)
	defer db.Close()

	err = db.Ping()
	checkErr(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/amigos", amigos)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/read", read)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", del)
	http.HandleFunc("/drop", drop)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":7777", nil)
	checkErr(err)
}

func index(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "at index")
	checkErr(err)
}

func amigos(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`SELECT * FROM Amigos;`)
	checkErr(err)
	defer rows.Close()

	// data to be used in query
	var s, name string
	s = "RETRIEVED RECORDS:\n"

	// query
	for rows.Next() {
		err = rows.Scan(&name)
		checkErr(err)
		s += name + "\n"
	}
	fmt.Fprintln(w, s)
}

func create(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`CREATE TABLE Customer (name VARCHAR(50));`)
	checkErr(err)
	defer stmt.Close()

	rd, err := stmt.Exec()
	checkErr(err)

	n, err := rd.RowsAffected()
	checkErr(err)

	fmt.Fprintln(w, "CREATED TABLE Customer", n)
}

func insert(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`INSERT INTO Customer VALUE ("Farkhan");`)
	checkErr(err)
	defer stmt.Close()

	rd, err := stmt.Exec()
	checkErr(err)

	n, err := rd.RowsAffected()
	checkErr(err)

	fmt.Fprintln(w, "INSERTED RECORD", n)
}

func read(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`SELECT * FROM Customer;`)
	checkErr(err)
	defer rows.Close()

	var name string
	for rows.Next() {
		err = rows.Scan(&name)
		checkErr(err)
		fmt.Fprintln(w, "RETRIEVED RECORD", name)
	}
}

func update(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`UPDATE Customer SET name="Hamzah" WHERE name="Farkhan";`)
	checkErr(err)
	defer stmt.Close()

	rd, err := stmt.Exec()
	checkErr(err)

	n, err := rd.RowsAffected()
	checkErr(err)

	fmt.Fprintln(w, "UPDATED RECORD", n)
}

func del(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`DELETE FROM Customer WHERE name="Hamzah";`)
	checkErr(err)
	defer stmt.Close()

	rd, err := stmt.Exec()
	checkErr(err)

	n, err := rd.RowsAffected()
	checkErr(err)

	fmt.Fprintln(w, "DELETED RECORD", n)
}

func drop(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`DROP TABLE Customer;`)
	checkErr(err)
	defer stmt.Close()

	rd, err := stmt.Exec()
	checkErr(err)

	_, err = rd.RowsAffected()
	checkErr(err)

	fmt.Fprintln(w, "DROPPED TABLE Customer")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
