package controller

import "html/template"

var (
	homeController home
	templates map[string]*template.Template
)

func init() {
	templates = PopulateTemplates()
}

// StartUp func
func StartUp() {
	homeController.registerRoutes()
}