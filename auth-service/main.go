package main

import (
  "fmt"
  "log"
  "net/http"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "This is the auth service : )")
}
func main() {

  http.HandleFunc("/", testHandler)
  err := http.ListenAndServe(":8080", nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
    return
  }
}
