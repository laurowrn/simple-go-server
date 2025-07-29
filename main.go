package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodGet {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }
        fmt.Fprintf(w, "Hello, this is a simple HTTP server")
    })

    fmt.Println("Server starting on http://0.0.0.0:8080")
    err := http.ListenAndServe("0.0.0.0:8080", nil)
    if err != nil {
        fmt.Printf("Server failed to start: %v\n", err)
    }
}

