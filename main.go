package main

import (
	"net/http"
	"template-module/templates"
)

func main() {
	// Serve static files (optional)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Route: full page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := templates.Index().Render(r.Context(), w)
		if err != nil {
			http.Error(w, "Failed to render template", http.StatusInternalServerError)
		}
	})

	// Route: HTMX partial using templ
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		/*err := templates.Hello("hello world").Render(r.Context(), w)
		if err != nil {
			http.Error(w, "Failed to render template", http.StatusInternalServerError)
		}*/
	})

	http.ListenAndServe(":8080", nil)
}
