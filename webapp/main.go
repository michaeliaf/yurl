package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/chayev/yurl/yurllib"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Listening on :8080...")
	http.HandleFunc("/", handler)
	http.HandleFunc("/results", viewResults)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type PageOutput struct {
	Content string
	URL     string
	Prefix  string
	Bundle  string
}

func handler(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("tpl/base.html")
	t.Execute(w, nil)
}

func viewResults(w http.ResponseWriter, r *http.Request) {

	url := r.URL.Query()["url"][0]

	prefix := r.URL.Query()["prefix"][0]

	bundle := r.URL.Query()["bundle"][0]

	var output []string

	if url == "" {
		output = []string{}
	}

	output = yurllib.CheckDomain(url, prefix, bundle, true)

	content := &PageOutput{URL: url, Prefix: prefix, Bundle: bundle}

	for _, item := range output {
		content.Content += item
	}

	t, _ := template.ParseFiles("tpl/results.html")
	t.Execute(w, &content)
}
