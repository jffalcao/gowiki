package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	title := "index"
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<H1>%s - Homepage Handler</H1><div>%s</div>"+
		"<a href=\"/edit\"> edit Handler</a><br>"+
		"<a href=\"/view/TestPage\"> view Handler</a>", p.Title, p.Body)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func main() {

	// Call the home page by typing http://localhost:8080
	http.HandleFunc("/", homeHandler)

	// Test the view page by typing http://localhost:8080/view/Testpage
	http.HandleFunc("/view/", viewHandler)

	// Test the edit page by typing http://localhost:8080/edit/
	http.HandleFunc("/edit/", editHandler)
	http.ListenAndServe(":8080", nil)
}
