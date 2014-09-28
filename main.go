package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Ping struct {
	Status  int `json:"status"`
	Results struct {
		Message string `json:"message"`
	} `json:"results"`
}

type ping_result struct {
	Message string `json:"message"`
}

func randString(n int) string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
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

	code := randString(8)

	db, err := sql.Open("mysql", "root@/test")
	if err != nil {
		fmt.Errorf("Error: %s", err)
	}

	npIns, err := db.Prepare("INSERT INTO np (code, title, body) VALUES(?, ?, ?)")
	if err != nil {
		fmt.Errorf("Insert Prepare Error: %s", err)
	}
	defer npIns.Close()

	_, err = npIns.Exec(code, title, body)

	if err != nil {
		fmt.Errorf("Insert Error: %s", err)
	}

	http.Redirect(w, r, "/np/"+code, http.StatusFound)
}

func npHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write(res)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	ping := Ping{http.StatusOK, ping_result{"ok"}}

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
