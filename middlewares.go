package express

import (
	"log"
	"net/http"
)

func LogPathAccess(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	log.Printf("%s `%s`\n", r.Method, path)
}
