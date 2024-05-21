package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getAiResponseHandler(w http.ResponseWriter, r *http.Request) {
    promptBuffer, err := io.ReadAll(r.Body)
    if err != nil {
        return
    }
    defer r.Body.Close()

    var prompt map[string]interface{}
    if err := json.Unmarshal(promptBuffer, &prompt); err != nil {
        return
    }

    question := prompt["content"].(string)

	fmt.Println("endpoint hit")
	
    answer, err := queryOllama(question)
    if err != nil {
        fmt.Fprintln(w, "Error querying Ollama:", err)
        return
    }
    fmt.Fprintln(w, "Ollama response:", answer)
}