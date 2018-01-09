package main

import (
  "fmt"
  "log"
  "os"

  // "net/http"
  // "goji.io"
  // "goji.io/pat"

  "github.com/joho/godotenv"
)

/*
func test(w http.ResponseWriter, r *http.Request) {
  name := pat.Param(r, "name")
  fmt.Fprintf(w, "Test, %s!", name)
}
*/

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  es_host := os.Getenv("ELASTICSEARCH_HOST00")
  fmt.Println(es_host)
/*
  mux := goji.NewMux()
  mux.HandleFunc(pat.Get("/test/:name"), test)

  http.ListenAndServe("localhost:8000", mux)
*/
}
