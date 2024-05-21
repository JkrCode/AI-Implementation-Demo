package main

import (
	"fmt"
	"net/http"
)

func main() {
    http.HandleFunc("/", getAiResponseHandler)
    fmt.Println("Starting server at port 8081")
    if err := http.ListenAndServe(":8081", nil); err != nil {
        fmt.Println("Failed to start server:", err)
        return
    }
}
