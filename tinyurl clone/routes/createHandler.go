package routes

import (
	"go.mongodb.org/mongo-driver/bson"
	"html/template"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
	"tinyurl-clone/db"
)

func CreateURL(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		ori := r.FormValue("url")
		if !strings.HasPrefix(ori, "http://") && !strings.HasPrefix(ori, "https://") {
			ori = "http://" + ori
		}
		_, err := url.ParseRequestURI(ori)
		if err != nil {
			t := template.New("index.html")
			t, _ = t.ParseFiles("static/index.html")
			vars := map[string]interface{}{
				"Created": true,
				"Content": "Invalid URL",
			}
			t.Execute(w, vars)
		} else {
			val := GenerateHash()
			for {
				_, found := db.Find(bson.D{{"hash", bson.D{{"$eq", val}}}})
				if found {
					val = GenerateHash()
					continue
				}
				db.Insert(bson.M{"hash": val, "url": ori})
				t := template.New("index.html")
				t, _ = t.ParseFiles("static/index.html")
				vars := map[string]interface{}{
					"Created": true,
					"Content": "Your url for " + ori + " is " + r.Host + "/" + val,
				}
				t.Execute(w, vars)
				break
			}
		}
	}
}

func GenerateHash() string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var out strings.Builder
	for i := 1; i <= 8; i++ {
		out.WriteString(string(letterBytes[r1.Intn(len(letterBytes))]))
	}
	return out.String()
}
