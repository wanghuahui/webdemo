package main

import (
	"html/template"
	"log"
	"net/http"
)

type loginController struct {
}

func (p *loginController) IndexAction(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/html/login/index.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
}
