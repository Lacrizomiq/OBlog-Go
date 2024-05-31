package controllers

import (
	"myapp/myapp/models"
	"myapp/myapp/views"
	"net/http"
	"path/filepath"
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
	tplPath := filepath.Join("views/templates", "article.gohtml")
	executeTemplate(w, tplPath, nil)
}
