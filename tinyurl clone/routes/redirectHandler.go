package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"tinyurl-clone/db"
)

func Redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hash := vars["hash"]
	result, found := db.Find(bson.D{{"hash", bson.D{{"$eq", hash}}}})
	if !found {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: Invalid URL!")
	} else {
		http.Redirect(w, r, result.URL, 302)
	}

}
