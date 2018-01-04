package main

import (
  "fmt"
  "net/http"

  "goji.io"
  "goji.io/pat"
)

func test(w http.ResponseWriter, r *http.Request) {
  name := pat.Param(r, "name")
  fmt.Fprintf(w, "Test, %s!", name)
}

func main() {
  mux := goji.NewMux()
  mux.HandleFunc(pat.Get("/test/:name"), test)

  http.ListenAndServe("localhost:8000", mux)
}
