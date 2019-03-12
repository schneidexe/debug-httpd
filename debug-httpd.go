package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "regexp"
    "strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Remote-Addr: %s\n", r.RemoteAddr)
    fmt.Fprintf(w, "Host: %s\n", r.Host)
    fmt.Fprintf(w, "Path: %s\n", r.URL.Path)
    fmt.Fprintf(w, "Query: %s\n", r.URL.RawQuery)
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
    fmt.Fprintf(w, "\nEnvironment:\n")
    for _, env := range os.Environ() {
        fmt.Fprintf(w, "\t%s\n", env)
    }

    if os.Getenv("UPSTREAM") != "" {
        resp, err := http.Get(os.Getenv("UPSTREAM"))
        if err != nil {

        }
        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)
        lineStarts := regexp.MustCompile(`(?m:^)`)
        tabbedBody := lineStarts.ReplaceAllString(string(body), "\t")
        fmt.Fprintf(w, "\nUpstream response:\n")
        fmt.Fprintf(w, "%s\n", tabbedBody)
    }
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}

