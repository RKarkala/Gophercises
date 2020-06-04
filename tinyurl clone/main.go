package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"tinyurl-clone/routes"
)

func main() {
	r := mux.NewRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	r.HandleFunc("/", routes.Home)
	r.HandleFunc("/create", routes.CreateURL)
	r.HandleFunc("/{hash}", routes.Redirect)
	http.Handle("/", r)
	http.ListenAndServe(":"+port, nil)
}
