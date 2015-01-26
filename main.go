package main 

import (
  "fmt"
  "net/http"

  "appengine"
  "appengine/user"
)

func init() {
  http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
  // fmt.Fprint(w, "Hello, World!")

  c := appengine.NewContext(r)
  u := user.Current(c) // if user has already logged in
  if u == nil {
    url, err := user.LoginURL(c, r.URL.String())
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    w.Header().Set("Location", url)
    w.WriteHeader(http.StatusFound)
    return
  }
  fmt.Fprintf(w, "Hello, %v!", u)
}
