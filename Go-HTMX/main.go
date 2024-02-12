package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Film struct{
	Title string
	Director string
}

func homeHandler(w http.ResponseWriter, r *http.Request){
	log.Println(r.URL.RawPath)

	tmpl := template.Must(template.ParseFiles("./templates/index.html"))

	films := map[string][]Film{
		"Films":{
			{Title: "The Witcher", Director: "John Stones"},
			{Title: "The Hunger Games", Director: "Susanne Collins"},
		},
	}

	tmpl.Execute(w, films)
}

func filmHandler(w http.ResponseWriter, r *http.Request){
	title := r.PostFormValue("title")
	director := r.PostFormValue("director")

	htmlStr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'> %s - %s </li>", title, director)

	tmpl, _ := template.New("t").Parse(htmlStr)
	tmpl.Execute(w, nil)
}

func main() {
	setupAPI()

	log.Printf("Server running on port:8000\n", )
	log.Fatal(http.ListenAndServe(":8000",nil))
}

func setupAPI(){
	http.Handle("/", http.FileServer(http.Dir("./templates")))
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/add_film/", filmHandler)
}