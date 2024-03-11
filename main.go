package main

import (
	"html/template"
	"net/http"
)

func main() {
    tmpl := template.Must(template.ParseFiles("template.html"))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data := struct {
            Title string
            Body  string
        }{
            Title: "My Page Title",
            Body:  "This is my page body.",
        }

        tmpl.Execute(w, data)
    })

    http.ListenAndServe(":8081", nil)
}
