package main

import (
	"net/http"

	"github.com/simenandre/templ-watch-bug/templates"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("internal/server/assets/public"))))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		landing := templates.LandingPage()
		landing.Render(r.Context(), w)
	})

	srv := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	srv.ListenAndServe()
}
