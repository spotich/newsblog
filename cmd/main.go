package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/spotich/newsblog/internal/pkg/newsmanager"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basePath   = filepath.Join(filepath.Dir(b), "../")
	mgr        newsmanager.Manager
)

func main() {
	os.Setenv("BASE_PATH", basePath)
	http.HandleFunc("/news", newsHandler)
	http.Handle("/public/img/", http.StripPrefix("/public/img/", http.FileServer(http.Dir(path.Join(basePath, "/public/img/")))))
	log.Println("server is listening at http://localhost:3000/news")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func newsHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles("../public/tmpl/news.html")
	if err != nil {
		log.Fatal(err)
	}
	configPath := fmt.Sprintf("%s/configs/database.json", basePath)

	err = mgr.Connect(configPath)
	if err != nil {
		log.Fatal(err)
	}
	news, err := mgr.GetNews()
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, news)
}
