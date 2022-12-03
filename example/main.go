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

func GetIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(express.GenericResponse{"message": "GET hello world"})
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

	app.Get("/", GetIndex)

	app.Get("/api/messages", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(express.GenericResponse{"message": "GET messages"})
	})

	app.Post("/api/messages", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(express.GenericResponse{"message": "POST messages"})
	})

	app.Put("/api/messages", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(express.GenericResponse{"message": "PUT messages"})
	})

	app.Patch("/api/messages", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(express.GenericResponse{"message": "PATCH messages"})
	})

	app.Delete("/api/messages", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(express.GenericResponse{"message": "DELETE messages"})
	})

	app.Listen(5200)
}
