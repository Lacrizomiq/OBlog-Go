package views

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Parse
func Parse(filepath string) (Template, error) {
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		return Template{}, fmt.Errorf("could not parse template: %w", err)
	}
	return Template{tpl}, nil
}

type Template struct {
	htmlTpl *template.Template
}

// Execute
func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html")
	err := t.htmlTpl.Execute(w, data)
	if err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
