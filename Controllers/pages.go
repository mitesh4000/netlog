package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"netLog/db"
	"strconv"
)

type PageData struct {
	Title   string
	Heading string
	Message string
}

func Home(w http.ResponseWriter, r *http.Request) {

	totalVisitors, err := db.GetTotalVisitors()
	if err != nil {
		log.Fatal("unable to count total visiters")
	}

	tmpl := template.Must(template.ParseFiles("./index.html"))
	if err != nil {
		http.Error(w, "Unable to load tempalte", http.StatusInternalServerError)
		return
	}

	data := PageData{
		Title:   "Server status",
		Heading: "yo server is up and running",
		Message: fmt.Sprintf("total %d visitors till day", &totalVisitors),
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}

}

func HelloHandler(w http.ResponseWriter, r *http.Request) {

	totalVisitors, err := db.GetTotalVisitors()
	if err != nil {
		log.Fatal("unable to count total visiters")
	}
	tmpl := `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>{{.Title}}</title>
  </head>
  <body>
    <h1>{{.Heading}}</h1>
    <p>{{.Message}}</p>
	<h6>hi man</h6>
  </body>
</html>`

	t, err := template.New("hello").Parse(tmpl)
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}

	data := PageData{
		Title:   "Server status",
		Heading: "yo server is up and running",
		Message: "Total " + strconv.Itoa(totalVisitors) + " visitors till date",
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
	}
}
