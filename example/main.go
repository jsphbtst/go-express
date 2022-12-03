package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jsphbtst/go-express"
)

func SampleCustomMiddleware(config express.Config) express.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/redirect" && r.Method == http.MethodGet {
			log.Println("REDIRECT URL IDK MAN")
		}
	}
}

func PostIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(express.GenericResponse{"message": "POST hello world"})
}

func main() {
	app := express.New()

	app.Use(express.Cors(
		express.Config{
			"origin": "*",
		},
	))
	app.Use(express.LogPathAccess(nil))
	app.Use(SampleCustomMiddleware(nil))

	app.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(express.GenericResponse{"message": "GET hello world"})
	})

	app.Post("/", PostIndex)

	app.Get("/messages", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(express.GenericResponse{"message": "GET messages"})
	})

	app.Listen(5200)
}
