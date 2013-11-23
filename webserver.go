package main

import (
    "fmt"
    "net/http"
)

type Server struct {}
func (h Server) ServeHTTP (
  w http.ResponseWriter,
  r *http.Request) {
  fmt.Fprintf(w, "Hi!, method=%s, url=%s, agent=%s\n", r.Method, r.URL.Path[4:], r.Header["User-Agent"])
}

type Server2 struct {}
func (h Server2) ServeHTTP (
  w http.ResponseWriter,
  r *http.Request) {
  fmt.Fprintf(w, "Server2 here\n")
}

// $ curl http://localhost:4000/path1/some/thing/here
// Hi!, method=GET, url=h2/some/thing/here, agent=[curl/7.28.1]
// $ curl http://localhost:4000/path2/some/thing/here
// Server2 here
func main() {
  var h Server
  var h2 Server2
  http.Handle("/path1/",h)
  http.Handle("/path2/",h2)
  http.ListenAndServe("localhost:4000",nil)
}
