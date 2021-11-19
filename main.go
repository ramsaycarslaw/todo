package main

import (
	"log"
	"net/http"
	"text/template"
)

// todo -- struct to store the todo and meta data
type todo struct {
	Title string
	Done  bool
}

// index -- the handler for the home page
func index(w http.ResponseWriter, r *http.Request) {
	// pasre all templates
	tpl, err := template.ParseFiles("public/todo.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	t2 := todo{Title: "Probability Hand In", Done: true}
	t3 := todo{Title: "Calculus Homework", Done: false}
	t4 := todo{Title: "Computer Systems Lecture", Done: false}
	t5 := todo{Title: "Gym - Chest Day", Done: true}
	t6 := todo{Title: "Algorithms Lab", Done: false}
	t7 := todo{Title: "Software Development Coursework", Done: false}
	t8 := todo{Title: "Algorithms Tutorial", Done: true}
	t9 := todo{Title: "Call Land Lord", Done: false}

	todos := []todo{t2, t3, t4, t5, t6, t7, t8, t9}

	err = tpl.Execute(w, todos)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// http routes
	http.HandleFunc("/", index)

	// Serve CSS and other static components
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
