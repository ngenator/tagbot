package main

import (
    "net/http"
    "fmt"
    "encoding/json"
    "io/ioutil"
    "net/url"
    "strings"
)

type WikiSearchResult struct {
  Title string `json:"title"`
  Size int `json:"size"`
  WordCount int `json:"wordcount"`
  Snippet string `json:"snippet"`
  Timestamp string `json:"timestamp"`
}

type WikiSearchQuery struct {
  SearchResults []WikiSearchResult `json:"search"`
}

type WikiSearchResponse struct {
  Query WikiSearchQuery `json:"query"`
}

func WikiSearch(term string) string {
  resp, err := http.Get("https://en.wikipedia.org/w/api.php?action=query&list=search&format=json&srsearch=" + url.QueryEscape(term))
  if err != nil {
    return "Error: [" + err.Error() + "]"
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)

  var result WikiSearchResponse
  if err := json.Unmarshal(body, &result); err != nil {
    return "Error: [" + err.Error() + "]"
  }
  t := &url.URL{Path: strings.Replace(result.Query.SearchResults[0].Title, " ", "_", -1)}
  return "https://en.wikipedia.org/wiki/" + t.String()
}

func execute(w http.ResponseWriter, r *http.Request) {
    args := r.URL.Query()["args"][0]
    fmt.Fprint(w, WikiSearch(args)) // send data to client side
}

func main() {
    http.HandleFunc("/execute", execute) // set router
    http.ListenAndServe(":80", nil) // set listen port
}
