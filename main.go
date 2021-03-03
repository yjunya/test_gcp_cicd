package main

import (
    "fmt"
    "net/http"
)

func main() {
    // ①
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, Build Trigger! Test error")
    })

    // ②
    http.ListenAndServe(":3000", nil)
}
