package routes

import (
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	t := template.New("index.html")
	t, _ = t.ParseFiles("static/index.html")
	vars := map[string]interface{}{
		"Created": false,
		"Content": "",
	}
	t.Execute(w, vars)

}
