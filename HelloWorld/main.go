package main

import (
  "fmt"
  "rsc.io/quote"
  "net/http"
)

func main() {
  // Request handler
   http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
      fmt.Fprintf(w, quote.Go(), r.URL.Path)
  }) 
  http.ListenAndServe(":80", nil) // Start server and listen for request on port 80
}
