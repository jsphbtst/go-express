# Go Express

This is currently an experiment to see how much of `express.js`, an HTTP server library for `Node.js` apps, I can copy. Yes, I know `go-fiber` exists, but I want to build this solely with net/http.

## Sample Code

Check out `example/main.go`, but in summary, this should feel much like Express.

```
app := express.New()

app.Use(express.Cors(
  express.Config{
    "origin": "*",
  },
))

app.Get("/", func(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(express.GenericResponse{"message": "GET hello world"})
})

app.Post("/", func(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(express.GenericResponse{"message": "POST hello world"})
})

app.Listen(5200)
```
