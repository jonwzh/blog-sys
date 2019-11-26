package controller

import (
	"html/template"

	"github.com/gorilla/sessions"
)

var (
	homeController home
	templates      map[string]*template.Template
	sessionName    string
	store          *sessions.CookieStore
)

func init() {
	templates = PopulateTemplates()
	store = sessions.NewCookieStore([]byte("a-top-secret")) // TODO: Move to config file
	sessionName = "blog-sys"
}

// StartUp func
func StartUp() {
	homeController.registerRoutes()
}