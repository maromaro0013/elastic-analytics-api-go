package main

import (
  "context"
  "fmt"
  "log"
  //"os"

  // "net/http"
  // "goji.io"
  // "goji.io/pat"

  "github.com/joho/godotenv"
  elastic "gopkg.in/olivere/elastic.v5"
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

  //es_host := os.Getenv("ELASTICSEARCH_HOST00")

  client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"), elastic.SetSniff(false))
  if err != nil {
    log.Fatalf("Connect failed: %v", err)
  }
  defer client.Stop()
  log.Print("Connected")

  res, err := client.ClusterState().Metric("version").Do(context.Background())
  if err != nil {
    panic(err)
  }
  fmt.Println("Cluster %q has nodes %d", res.ClusterName, res.Version)

  names, err := client.IndexNames()
  for _, name := range names {
    fmt.Println(name)
  }

/*
  mux := goji.NewMux()
  mux.HandleFunc(pat.Get("/test/:name"), test)

  http.ListenAndServe("localhost:8000", mux)
*/
}
