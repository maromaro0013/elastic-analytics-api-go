package main

import (
  "context"
  "fmt"
  "log"
  "encoding/json"
  //"os"

  // "net/http"
  // "goji.io"
  // "goji.io/pat"

  "github.com/joho/godotenv"
  "gopkg.in/olivere/elastic.v5"
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
  fmt.Printf("Cluster %q has version %d\n", res.ClusterName, res.Version)

  //query := elastic.NewBoolQuery()
  //query = query.Must(elastic.NewRangeQuery("time_iso8601").Gte("2018-01-10T00:00:00Z").Lte("2018-01-11T00:00:00Z"))

  //query := elastic.NewRangeQuery("time_iso8601").Gte("2017-08-31T15:00:00Z").Lte("2017-09-30T14:59:59Z")
  query := elastic.NewMatchAllQuery()

/*
  termQuery := elastic.NewTermQuery("test", "test")
  searchResult, err := client.Search().
      Index("axs_corporate").
      Type("access_log").
      Query(termQuery).
      From(0).Size(10).
      Pretty(true).
      Do(context.Background())
  if err != nil {
      panic(err)
  }
*/
  src, err := query.Source()
  if err != nil {
  panic(err)
  }
  data, err := json.Marshal(src)
  if err != nil {
  panic(err)
  }
  s := string(data)
  fmt.Println(s)

  search := client.Search().Index("axs_corporate").Type("access_log")
  //search = search.Query(query)
  searchResult, err := search.Do(context.Background())
  if err != nil {
    panic(err)
  }

  fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)
  fmt.Printf("%d\n", searchResult.Hits.TotalHits)

/*
  mux := goji.NewMux()
  mux.HandleFunc(pat.Get("/test/:name"), test)

  http.ListenAndServe("localhost:8000", mux)
*/
}
