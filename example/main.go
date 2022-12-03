package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jsphbtst/go-express"
)

func main() {
	app := express.New()

	app.Get("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s `%s`", r.Method, r.URL)

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(express.GenericResponse{"message": "GET hello world"})
	})

	app.Get("/messages", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s `%s`", r.Method, r.URL)

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(express.GenericResponse{"message": "GET messages"})
	})

	app.Post("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s `%s`", r.Method, r.URL)

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(express.GenericResponse{"message": "POST hello world"})
	})

	app.Listen(5200)
}
