package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

// CarregarTemplates insere os templates html na varia templates
func CarregarTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
	templates = template.Must(templates.ParseGlob("views/templates/*.html"))
}

// ExecutarTemplate uma pagina Html
func ExecutarTemplate(w http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(w, template, dados)
}
