package main

import (
    "fmt"
    "net/http"
    "os"
    "strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Context-Path: /%s\n", r.URL.Path[1:])
    fmt.Fprintf(w, "Host: %s\n", r.Host)
    fmt.Fprintf(w, "Remote-Addr: %s\n", r.RemoteAddr)
    fmt.Fprintf(w, "Headers:\n")
    for k,vs := range r.Header {
        fmt.Fprintf(w, "\t%s:\n", k)
        for _, v := range vs {
            if (k == "Cookie") {
                v = strings.Join(strings.Split(v, "; "), "\n\t\t")
            }
            fmt.Fprintf(w, "\t\t%s\n", v)
        }
    }
    fmt.Fprintf(w, "Environment:\n")
    for _, env := range os.Environ() {
        fmt.Fprintf(w, "\t\t%s\n", env)
    }
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}

