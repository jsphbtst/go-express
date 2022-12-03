package express

import (
	"log"
	"net/http"
)

func LogPathAccess(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	log.Printf("%s `%s`\n", r.Method, path)
}

// TODO: allow for origin to be changed lol
func Cors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}
