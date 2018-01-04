package main

import (
	"fmt"
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
		"<a href=\"/edit/testpage\"> edit Handler</a><br>"+
		"<a href=\"/view/testpage\"> view Handler</a>", p.Title, p.Body)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<H1>%s - View Handler</H1><div>%s</div><br>", p.Title, p.Body)
	fmt.Fprintf(w, "<p><a href=\"/\"> Home</a></p>")
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	fmt.Fprintf(w, "<h1>Editing %s</h1>"+
		"<form action=\"/save/%s\" method=\"POST\">"+
		"<textarea name=\"body\">%s</textarea><br>"+
		"<input type=\"submit\" value=\"Save\">"+
		"</form>",
		p.Title, p.Title, p.Body)
	fmt.Fprintf(w, "<p><a href=\"/\"> Home</a></p>")
}

func main() {
	// part1 at beginning
	// p1 := &Page{Title: "TestPage", Body: []byte("This is a sample page.")}
	// p1.save()
	// p2, _ := loadPage("TestPage")
	// fmt.Println(string(p2.Body))

	// Part2 - starts at Introducing the net/http package (an interlude)

	// This will dispaly the content of TestPage.txt
	//Part3 - Starts at Editing Pages

	// Call the home page by typing http://localhost:8080
	http.HandleFunc("/", homeHandler)

	// Test the view page by typing http://localhost:8080/view/Testpage
	http.HandleFunc("/view/", viewHandler)

	// Test the edit page by typing http://localhost:8080/edit/Testpage
	http.HandleFunc("/edit/", editHandler)
	http.ListenAndServe(":8080", nil)
}
