package main

import (
      "fmt"
      "net/http"
      "bytes"
      "encoding/json"
      "flag"
    )

var printHeaders = false

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("---- BEGIN HTTP REQUEST ----")
  fmt.Println("Method: " + r.Method)
  fmt.Println("Path: " + r.URL.Path)
  if (printHeaders == true) { 
    jsonString, _ := json.MarshalIndent(r.Header, "", "  ")
    fmt.Println("Headers: \n" + string(jsonString[:]))
  }
  buffer := new(bytes.Buffer)
  buffer.ReadFrom(r.Body)
  fmt.Println("Body: " + buffer.String())
  fmt.Println("---- END HTTP REQUEST ----\n")
  fmt.Fprintf(w, "200 OK")
}

func main() {
  flag.BoolVar(&printHeaders, "headers", false, "should print headers")
  flag.Parse()

  http.HandleFunc("/", handler)
  http.ListenAndServe(":3000", nil)
}
