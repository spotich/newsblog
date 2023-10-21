package main

import (
	"html/template"
	"log"
	"net/http"
)

type news struct {
	Title string
	Body  string
}

func main() {
	http.HandleFunc("/news", newsHandler)
	log.Println("server is listening at http://localhost:3000/news")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func newsHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles("../public/news.html")
	if err != nil {
		log.Fatal(err)
	}
	data := getNews()
	tmpl.Execute(w, data)
}

func getNews() []news {
	return []news{
		{"First Title", "First Body"},
		{"Second Title", "Second Body"},
		{"Third Title", "Third Body"},
	}
}
