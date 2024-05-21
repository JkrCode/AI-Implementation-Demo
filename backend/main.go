package main

import (
	"fmt"
	"net/http"

	"github.com/rs/cors"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/api", getAiResponseHandler)

    // Enable CORS
    c := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://127.0.0.1:5173"}, // Change this to the origin of your React app
        AllowCredentials: true,
        AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
        AllowedHeaders:   []string{"Content-Type"},
    })

    handler := c.Handler(mux)
    fmt.Println("Starting server at port 8081")
    if err := http.ListenAndServe(":8081", handler); err != nil {
        fmt.Println("Failed to start server:", err)
        return
    }
}
