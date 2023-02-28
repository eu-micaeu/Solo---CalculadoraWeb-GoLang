package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id    int
	Name  string
	Email string
}

func main() {
	db, err := sql.Open("mysql", "root:M1ca3l13*@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	http.HandleFunc("/insert", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		email := r.FormValue("email")

		_, err := db.Exec("INSERT INTO users(name, email) VALUES (?, ?)", name, email)
		if err != nil {
			panic(err.Error())
		}

		fmt.Fprintln(w, "Usuário inserido com sucesso!")
	})

	fmt.Println("Servindo em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
