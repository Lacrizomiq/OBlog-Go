package main

import (
	"log"
	"myapp/myapp/controllers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	fs := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	r.Get("/", controllers.HomeHandler)
	r.Get("/article/{id}", controllers.ArticlePageHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", http.StatusNotFound)
	})

	log.Println("Starting server on :3000")
	http.ListenAndServe(":3000", r)
}
