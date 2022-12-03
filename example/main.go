package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jsphbtst/go-express"
)

func SampleCustomMiddleware(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/redirect" && r.Method == http.MethodGet {
		log.Println("REDIRECT URL IDK MAN")
	}
}

func main() {
	app := express.New()
	app.Use(express.Cors)
	app.Use(express.LogPathAccess)
	app.Use(SampleCustomMiddleware)

	app.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(express.GenericResponse{"message": "GET hello world"})
	})

	app.Get("/messages", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(express.GenericResponse{"message": "GET messages"})
	})

	app.Post("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(express.GenericResponse{"message": "POST hello world"})
	})

	app.Listen(5200)
}
