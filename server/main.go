package main

import (
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

type PageData struct {
	Success bool
	Url     string
}

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// provide assets
	assetsDir := filepath.Join(wd, "/internal/assets")
	fs := http.FileServer(http.Dir(assetsDir))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// create template
	templateDir := filepath.Join(wd, "/internal/templates/index.html")
	tmpl := template.Must(template.ParseFiles(templateDir))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			tmpl.Execute(w, nil)
			return
		}

		url := r.FormValue("url")

		data := PageData{
			Success: true,
			Url:     url,
		}
		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":8080", nil)
}
