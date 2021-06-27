package main

import (
	"../chat/trace"
	"flag"
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
		t.template = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	if err := t.template.Execute(w, r); err != nil {
		log.Fatal("Execute", err)
	}
}

// run under commands, if you want to run this program
// # go build -o chat
// # ./chat -addr=":3000"
func main() {
	var address = flag.String("addr", ":8080", "Application Address")
	flag.Parse()
	r := newRoom()
	r.tracer = trace.New(os.Stdout)
	http.Handle("/", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login",&templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/",loginHandler)
	http.Handle("/room", r)
	go r.run()
	log.Println("Start Web Sever. Port:", *address)
	if err := http.ListenAndServe(*address, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
	// Access http://localhost:8080/
}
