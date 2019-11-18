package main

import (
	"net/http"

	"github.com/jonwzh/blog-sys/controller"
)

func main() {
	controller.StartUp()
	http.ListenAndServe(":8080", nil)
}