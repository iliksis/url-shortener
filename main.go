package main

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"
)

type PageData struct {
	Success bool
	Url     string
}

//go:embed internal/assets
var assets embed.FS

//go:embed internal/templates/index.html
var tmplSource string

func main() {
	// provide assets
	fSys, err := fs.Sub(assets, "internal/assets")
	if err != nil {
		panic(err)
	}
	fs := http.FileServer(http.FS(fSys))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// create template
	tmpl := template.Must(template.New("index.html").Parse(tmplSource))

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
