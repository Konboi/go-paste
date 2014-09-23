package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Ping struct {
	Status int    `json:"status"`
	Result string `json:"result"`
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("view/form.html")
	if err != nil {
		fmt.Errorf("Error: %s", err)
	}
	t.Execute(w, nil)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	title := r.PostFormValue("title")
	body := r.PostFormValue("body")

	db, err := sql.Open("mysql", "root@/nopaste")
	if err != nil {
		fmt.Errorf("Error: %s", err)
	}
	rows, _ := db.Query("SHOW TABLES")
	log.Printf("%T", rows.Columns)
	log.Printf("title: %s \n body:%s", title, body)
}

func npHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("view/np.html")
	if err != nil {
		fmt.Errorf("Error: %s", err)
	}
	t.Execute(w, nil)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	ping := Ping{http.StatusOK, "ok"}

	res, err := json.Marshal(ping)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func main() {
	http.HandleFunc("/", formHandler)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/np", npHandler)
	http.ListenAndServe(":8000", nil)
}
