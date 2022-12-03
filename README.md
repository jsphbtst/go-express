# Go Express

This is currently an experiment to see how much of `express.js`, an HTTP server library for `Node.js` apps, I can copy. Yes, I know `go-fiber` exists, but I want to build this solely with net/http.

Currently supported HTTP methods are `GET`, `POST`, `PUT`, `PATCH`, and `DELETE`.

## Sample Code

Check out `example/main.go`, but in summary, this should feel much like Express.

```
func GetIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(express.GenericResponse{"message": "POST hello world"})
}

...

app := express.New()

app.Use(express.Cors(
  express.Config{
    "origin": "*",
  },
))

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
```
