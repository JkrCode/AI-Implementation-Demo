package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)


func queryOllama(question string) (string, error) {
    // Create request payload with conversation context
    payload := map[string]interface{}{
        "model": "llama3",
        "messages": []map[string]string{
            {
                "role":    "user",
                "content": question,
            },
        },
        "stream": false,
    }
    jsonData, err := json.Marshal(payload)
    if err != nil {
        return "", err
    }

    // Prepare the request to the Ollama chat endpoint
    response, err := http.Post("http://localhost:11434/api/chat", "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        return "", err
    }
    defer response.Body.Close()

    // Read the response body
    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return "", err
    }

    // Parse the JSON response
    var ollamaResponse map[string]interface{}
    if err := json.Unmarshal(body, &ollamaResponse); err != nil {
        return "", err
    }

    // Extract the response message content from the Ollama response
    message, ok := ollamaResponse["message"].(map[string]interface{})
    if !ok {
        return "", fmt.Errorf("Invalid response from Ollama")
    }

    content, ok := message["content"].(string)
    if !ok {
        return "", fmt.Errorf("Invalid response content from Ollama")
    }

    return content, nil
}




func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
    question := "What is the capital of France?"
	fmt.Println("endpoint hit")
    answer, err := queryOllama(question)
	fmt.Println(answer)
    if err != nil {
        fmt.Fprintln(w, "Error querying Ollama:", err)
        return
    }
    fmt.Fprintln(w, "Ollama response:", answer)
}

func main() {
    http.HandleFunc("/", helloWorldHandler)
    fmt.Println("Starting server at port 8081")
    if err := http.ListenAndServe(":8081", nil); err != nil {
        fmt.Println("Failed to start server:", err)
        return
    }
}
