package express

import (
	"log"
	"net/http"
)

func LogPathAccess(config Config) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		log.Printf("%s `%s`\n", r.Method, path)
	}
}

func Cors(config Config) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		origin, ok := config["origin"]
		if !ok {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		} else {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}
	}
}
