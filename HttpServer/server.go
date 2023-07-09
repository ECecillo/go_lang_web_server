package main
// TODO 1: Process dynamic requests: Process incoming requests from users who browse
// The website, log into their accounts or post images.

// TODO 2: Serve static assets: Serve JS, CSS and images to browsers to create a 
// dynamic experience for the user.

// TODO 3: Listen on a specific port. 

import (
  "fmt"
  "net/http"
)

func main() {
  // Add a handler on root path
  http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello there Mr Kenobi")
  }) 

  // For file serving we use http.FileServer
  // Here we create a fileDescriptor instance that we can then use with http.Handle
  // To retrieve the file requested by a client.
  fs := http.FileServer(http.Dir("static/"))

  // Tell our web server that for any ressource requested on /static/ use the 
  // fs instance to read the file.
  http.Handle("/static/", http.StripPrefix("/static/", fs))

  // Listen on port 8080 now
  http.ListenAndServe(":80", nil)
}

