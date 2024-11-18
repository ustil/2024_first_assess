package main

import (
	"html/template"
	"log"
	"net/http"
)

type AAA struct {
	AAA string
}

func main() {
	tmpl, err := template.ParseFiles("001/a/index.html")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		vars := AAA{AAA: "你好"}
		err = tmpl.Execute(w, vars)
		if err != nil {
			log.Fatal(err)
		}
	})

	log.Println("服务器运行在 http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}