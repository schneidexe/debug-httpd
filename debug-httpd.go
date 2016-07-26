package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Context-Path: /%s\n", r.URL.Path[1:])
    fmt.Fprintf(w, "Host: %s\n", r.Host)
    fmt.Fprintf(w, "Headers:\n")
    for k,vs := range r.Header {
        fmt.Fprintf(w, "\t%s:\n", k)
        for _, v := range vs {
            fmt.Fprintf(w, "\t\t%s\n", v)
        }
    }
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}

