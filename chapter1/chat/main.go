package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"text/template"
)

type templateHandler struct {
	once     sync.Once
	filename string
	template *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		path, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		fmt.Println(path)
		// current directory is chapter1
		t.template = template.Must(template.ParseFiles(filepath.Join("chat/templates", t.filename)))
	})
	if err := t.template.Execute(w, nil); err != nil {
		log.Fatal("Execute",err)
	}
}

func main() {
	http.Handle("/", &templateHandler{filename: "chat.html"})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
	// Access http://localhost:8080/
}
