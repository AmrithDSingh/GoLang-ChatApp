package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type Welcome struct {
	Sale string
	Time string
}

func main() {
	welcome := Welcome{"ask away", time.Now().Format(time.Stamp)}
	template := template.Must(template.ParseFiles("static/let-us-see.html"))
	fmt.Print("test")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if err := template.ExecuteTemplate(w, "let-us-see.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	fmt.Println(http.ListenAndServe(":8000", nil))
}
