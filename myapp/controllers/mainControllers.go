package controllers

import (
	"myapp/myapp/models"
	"myapp/myapp/views"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func executeTemplate(w http.ResponseWriter, filename string, data interface{}) {
	tpl, err := views.Parse(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tpl.Execute(w, data)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("views/templates", "index.gohtml")

	articles, err := models.LoadArticles("models/articles.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	executeTemplate(w, tplPath, articles)
}

func ArticlePageHandler(w http.ResponseWriter, r *http.Request) {
	articleIDStr := chi.URLParam(r, "id")
	articleID, err := strconv.Atoi(articleIDStr)
	if err != nil {
		http.Error(w, "Invalid article ID", http.StatusBadRequest)
		return
	}

	articles, err := models.LoadArticles("models/articles.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var article models.Article
	for _, a := range articles {
		if a.ID == articleID {
			article = a
			break
		}
	}

	if article.ID == 0 {
		http.Error(w, "Article not found", http.StatusNotFound)
		return
	}

	tplPath := filepath.Join("views/templates", "article.gohtml")
	executeTemplate(w, tplPath, article)
}
