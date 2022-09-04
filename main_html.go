package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

// a struct for user validation

type userdetails struct {
	Username string
	Password string
}

// a variable to access our template directory

var tpl *template.Template

// this function initialize the path

func init() {
	tpl = template.Must(template.ParseGlob("template/*"))
}

// this function does the login logic
func loginLogic(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	user := new(userdetails)
	decoder := schema.NewDecoder()

	// decode the url.values which is map[string] string to the struct

	decodeErr := decoder.Decode(user, r.PostForm)

	if decodeErr != nil {
		log.Print("error mapping parsed form data to a struct :", decodeErr)
	}

	fmt.Fprintln(w, "Username:", user.Username)

}

//this  handler-function renders the pages
func loginRendering(w http.ResponseWriter, r *http.Request) {

	fmt.Println("method:", r.Method)

	err := tpl.ExecuteTemplate(w, "Login.html", nil)

	if err != nil {
		log.Fatal("cant execute html :", err)
		return
	}

}

func main() {

	// // using gorilla mux to route the url to our handler
	r := mux.NewRouter()

	r.HandleFunc("/", loginRendering)
	r.HandleFunc("/login", loginLogic)

	// creating a file server for the static files

	fileServer := http.FileServer(http.Dir("static"))

	// this removes the /static/ from url query
	r.PathPrefix("/").Handler(http.StripPrefix("/static/", fileServer))

	http.ListenAndServe(":8080", r)
}
