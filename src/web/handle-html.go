package web

import (
	"html/template"
	"net/http"
)

func HandleHTML(res http.ResponseWriter, req *http.Request) {
	indexTmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	indexTmpl.Execute(res, map[string]interface{}{
		"videoSrc": "19a91307-d5ab-47bd-a9ee-8f08abcd1229",
	})
}
